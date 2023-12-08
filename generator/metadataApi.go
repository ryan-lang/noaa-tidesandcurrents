package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GenerateMetadataApiClient(def MetadataApiDefinition) error {

	modelsPath := "./client/metadataApi/model.gen.go"
	err := GenerateModels(modelsPath, "metadataApi", def.Model)
	if err != nil {
		return err
	}

	for _, resourceDef := range def.StationResources {
		err = GenerateMetadataApiStationResource(&resourceDef)
		if err != nil {
			return err
		}
	}

	err = GenerateMetadataApiTests(def)
	if err != nil {
		return err
	}

	return nil
}

func GenerateMetadataApiStationResource(resource *StationResourceDefinition) error {

	// open file for writing; overrite if exists
	path := fmt.Sprintf("./client/metadataApi/%s.gen.go", resource.ResourceID)
	_ = os.MkdirAll(filepath.Dir(path), 0700)
	log.Printf("generating metadata api file: %s", path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	// write package
	f.WriteString("package metadataApi\n\n")

	// write gen warning
	f.WriteString("// THIS FILE IS GENERATED. DO NOT EDIT.\n\n")

	// write imports
	f.WriteString("import (\n")
	f.WriteString("\t\"context\"\n")
	f.WriteString("\t\"encoding/json\"\n")
	f.WriteString("\t\"fmt\"\n")
	f.WriteString("\t\"log\"\n")
	f.WriteString("\t\"github.com/pkg/errors\"\n")
	if resource.RequestType != "" {
		f.WriteString("\t\"github.com/google/go-querystring/query\"\n")
	}
	f.WriteString(")\n\n")

	// write the method signature
	if resource.RequestType != "" {
		f.WriteString(fmt.Sprintf("func (c *StationRequest) %s(ctx context.Context, req *%s) (*%s, error) {\n\n", resource.Name, resource.RequestType, resource.ResponseType))
	} else {
		f.WriteString(fmt.Sprintf("func (c *StationRequest) %s(ctx context.Context) (*%s, error) {\n\n", resource.Name, resource.ResponseType))
	}

	// check the fetched metadata to see if the resource is available
	if len(resource.Availability) > 0 {
		var availCheck []string
		for _, avail := range resource.Availability {
			availCheck = append(availCheck, fmt.Sprintf("stationType == \"%s\"", avail))
		}
		f.WriteString("\t// check the fetched metadata to see if the resource is available\n")
		f.WriteString("\tif c.Metadata != nil {\n")
		f.WriteString("\t\tvar isResourceAvailable bool\n")
		f.WriteString("\t\tfor _, stationType := range c.Metadata.StationTypes() {\n")
		f.WriteString(fmt.Sprintf("\t\t\tif %s {\n", strings.Join(availCheck, " || ")))
		f.WriteString("\t\t\t\tisResourceAvailable = true\n")
		f.WriteString("\t\t\t\tbreak\n")
		f.WriteString("\t\t\t}\n")
		f.WriteString("\t\t}\n")
		f.WriteString("\t\tif !isResourceAvailable {\n")
		f.WriteString(fmt.Sprintf("\t\t\tlog.Printf(\"fetched metadata incidicates %s is not available for station %%s\", c.StationID)\n", resource.Name))
		f.WriteString("\t\t}\n")
		f.WriteString("\t} else {\n")
		f.WriteString("\t\tif c.client.Verbose {\n")
		f.WriteString(fmt.Sprintf("\t\t\tlog.Printf(\"availability of %s for station %%s is unknown. call FetchMetadata() first. trying anyway...\", c.StationID)\n", resource.Name))
		f.WriteString("\t\t}\n")
		f.WriteString("\t}\n\n")
	}

	// validate the request
	if resource.RequestType != "" {
		f.WriteString("\t// validate the request\n")
		f.WriteString("\tif err := req.Validate(); err != nil {\n")
		f.WriteString("\t\treturn nil, err\n")
		f.WriteString("\t}\n\n")

		// build the params
		f.WriteString("\t// build the params\n")
		f.WriteString("\tparams, _ := query.Values(req)\n")
	}

	var urlPath string
	if resource.ResourceID == "metadata" {
		urlPath = "/stations/%s.json"
	} else {
		urlPath = fmt.Sprintf("/stations/%%s/%s.json", resource.ResourceID)
	}

	// make the request
	f.WriteString("\t// make the request\n")
	if resource.RequestType != "" {
		f.WriteString(fmt.Sprintf("\trespBody, err := c.client.httpGet(ctx, fmt.Sprintf(\"%s\", c.StationID), params)\n", urlPath))
	} else {
		f.WriteString(fmt.Sprintf("\trespBody, err := c.client.httpGet(ctx, fmt.Sprintf(\"%s\", c.StationID), nil)\n", urlPath))
	}
	f.WriteString("\tif err != nil {\n")
	f.WriteString("\t\treturn nil, err\n")
	f.WriteString("\t}\n\n")

	// parse the response
	f.WriteString("\t// parse the response\n")
	f.WriteString("\tvar resp " + resource.ResponseType + "\n")
	f.WriteString("\terr = json.Unmarshal(respBody, &resp)\n")
	f.WriteString("\tif err != nil {\n")
	f.WriteString("\t\treturn nil, errors.Wrap(err, \"failed to parse response\")\n")
	f.WriteString("\t}\n\n")

	// return the response
	f.WriteString("\treturn &resp, nil\n")
	f.WriteString("}\n\n")

	// write the method signature
	f.WriteString(fmt.Sprintf("func (c *StationsRequest) %s(ctx context.Context) ([]*%s, error) {\n\n", resource.Name, resource.ResponseType))
	f.WriteString("\t\t// TODO: not yet implemented\n")
	f.WriteString("\t\treturn nil, nil\n")
	f.WriteString("}\n\n")

	return nil
}

func GenerateMetadataApiTests(def MetadataApiDefinition) error {

	// open file for writing; overrite if exists
	path := "./client/metadataApi/client_test.go"
	_ = os.MkdirAll(filepath.Dir(path), 0700)
	log.Printf("generating data api test file: %s", path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	// write package
	f.WriteString("package metadataApi_test\n\n")

	// write gen warning
	f.WriteString("// THIS FILE IS GENERATED. DO NOT EDIT.\n\n")

	// write imports
	f.WriteString("import (\n")
	f.WriteString("\t\"context\"\n")
	f.WriteString("\t\"testing\"\n")
	//f.WriteString("\t\"time\"\n")
	f.WriteString("\t\"encoding/json\"\n")
	f.WriteString("\t\"fmt\"\n")
	f.WriteString("\t\"github.com/stretchr/testify/assert\"\n")
	f.WriteString("\t\"github.com/ryan-lang/noaa-tidesandcurrents/client/metadataApi\"\n")
	f.WriteString(")\n\n")

	// write the tests
	for _, resourceDef := range def.StationResources {
		f.WriteString(fmt.Sprintf("func Test%s(t *testing.T) {\n", resourceDef.Name))
		f.WriteString("\tc := metadataApi.NewClient(true, \"test\")\n")
		if resourceDef.TestStation == "" {
			resourceDef.TestStation = "9447130"
		}
		f.WriteString(fmt.Sprintf("\treq := metadataApi.NewStationRequest(c, \"%s\")\n", resourceDef.TestStation))
		f.WriteString("\tctx := context.Background()\n")

		// execute the request
		if resourceDef.RequestType != "" {
			f.WriteString(fmt.Sprintf("\tres, err := req.%s(ctx, &metadataApi.%s{})\n", resourceDef.Name, resourceDef.RequestType))
		} else {
			f.WriteString(fmt.Sprintf("\tres, err := req.%s(ctx)\n", resourceDef.Name))
		}
		f.WriteString("\tassert.NoError(t, err)\n")

		// write the response to json and print it
		f.WriteString("jsonBytes, _ := json.MarshalIndent(res, \"\", \"  \")\n")
		f.WriteString(fmt.Sprintf("fmt.Printf(\"%s response: %%s\\n\", jsonBytes)\n\n", resourceDef.Name))

		f.WriteString("}\n\n")
	}

	// gofmt the file
	err = goFmtFile(path)
	if err != nil {
		return err
	}

	return nil
}
