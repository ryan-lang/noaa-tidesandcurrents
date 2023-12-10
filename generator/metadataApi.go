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

	err = GenerateMetadataApiCli(def)
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

func GenerateMetadataApiCli(def MetadataApiDefinition) error {

	// write a get-data command
	path := "./cmd/noaatc/root/getMetadata/getMetadata.go"
	_ = os.MkdirAll(filepath.Dir(path), 0700)
	log.Printf("generating getMetadata: %s", path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	// write package
	f.WriteString("package getMetadata\n\n")

	// write gen warning
	f.WriteString("// THIS FILE IS GENERATED. DO NOT EDIT.\n\n")

	// write imports
	f.WriteString("import (\n")
	f.WriteString("\t\"github.com/spf13/cobra\"\n")
	for _, resource := range def.StationResources {
		f.WriteString(fmt.Sprintf("\t\"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getMetadata/%s\"\n", ToCamelCase(resource.ResourceID)))
	}
	f.WriteString(")\n\n")

	// write cmd
	f.WriteString("var GetMetadataCmd = &cobra.Command{\n")
	f.WriteString("\tUse:   \"getMetadata\",\n")
	f.WriteString("\tShort: \"Get data from NOAA CO-OPS Metadata API\",\n")
	f.WriteString("\tLong: `Get data from NOAA CO-OPS Metadata API`,\n")
	f.WriteString("}\n\n")

	// write init func
	f.WriteString("func init() {\n")
	for _, resource := range def.StationResources {
		f.WriteString(fmt.Sprintf("\tGetMetadataCmd.AddCommand(%s.%sCmd)\n", ToCamelCase(resource.ResourceID), ToUpperCamelCase(resource.ResourceID)))
	}
	f.WriteString("}\n\n")

	for _, resource := range def.StationResources {

		// open file for writing; overrite if exists
		path := fmt.Sprintf("./cmd/noaatc/root/getMetadata/%s/%s.go", ToCamelCase(resource.ResourceID), ToCamelCase(resource.ResourceID))
		_ = os.MkdirAll(filepath.Dir(path), 0700)
		log.Printf("generating data api test file: %s", path)
		f, err := os.Create(path)
		if err != nil {
			return err
		}

		// write package
		f.WriteString(fmt.Sprintf("package %s\n\n", ToCamelCase(resource.ResourceID)))

		// write gen warning
		f.WriteString("// THIS FILE IS GENERATED. DO NOT EDIT.\n\n")

		reqModel := def.GetModel(resource.RequestType)

		var hasDateParam, hasMultDateParams bool
		for _, fieldDef := range reqModel.Fields {
			if fieldDef.Type == "DateParam" {
				if hasDateParam {
					hasMultDateParams = true
					break
				}
				hasDateParam = true
			}
		}

		// write imports
		f.WriteString("import (\n")
		f.WriteString("\t\"github.com/spf13/cobra\"\n")
		f.WriteString("\t\"github.com/ryan-lang/noaa-tidesandcurrents/client/metadataApi\"\n")
		if hasDateParam {
			f.WriteString("\t\"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getMetadata/util\"\n")
		}
		f.WriteString("\t\"context\"\n")
		f.WriteString("\t\"encoding/json\"\n")
		f.WriteString("\t\"fmt\"\n")
		f.WriteString("\t\"log\"\n")
		f.WriteString(")\n\n")

		// write variables
		f.WriteString("var (\n")
		f.WriteString("\tstationId string\n")
		for _, fieldDef := range reqModel.Fields {
			switch fieldDef.Type {
			case "DateParam":
				f.WriteString(fmt.Sprintf("\t%sBeginDate string\n", ToCamelCase(fieldDef.Name)))
				f.WriteString(fmt.Sprintf("\t%sEndDate string\n", ToCamelCase(fieldDef.Name)))
				f.WriteString(fmt.Sprintf("\t%sRangeHours string\n", ToCamelCase(fieldDef.Name)))
				f.WriteString(fmt.Sprintf("\t%sRelative string\n", ToCamelCase(fieldDef.Name)))
			case "IntervalParam":
				f.WriteString(fmt.Sprintf("\t%s string\n", ToCamelCase(fieldDef.Name)))
			case "VelocityTypeParam":
			default:
				f.WriteString(fmt.Sprintf("\t%s %s\n", ToCamelCase(fieldDef.Name), fieldDef.Type))
			}
		}
		f.WriteString(")\n\n")

		// write cmd
		desc := fmt.Sprintf("Get %s data", resource.ResourceID)
		f.WriteString(fmt.Sprintf("var %sCmd = &cobra.Command{\n", ToUpperCamelCase(resource.ResourceID)))
		f.WriteString(fmt.Sprintf("\tUse:   \"%s\",\n", ToKebabCase(resource.ResourceID)))
		f.WriteString(fmt.Sprintf("\tShort: \"%s\",\n", desc))
		f.WriteString(fmt.Sprintf("\tLong: `%s`,\n", desc))
		f.WriteString(fmt.Sprintf("\tRun: func(cmd *cobra.Command, args []string) {\n"))
		f.WriteString(fmt.Sprintf("\t\tverbose, _ := cmd.Flags().GetBool(\"verbose\")\n"))
		f.WriteString(fmt.Sprintf("\t\tc := metadataApi.NewClient(verbose, \"github.com/ryan-lang/noaa-tidesandcurrents\")\n"))
		f.WriteString(fmt.Sprintf("\t\treq := metadataApi.NewStationRequest(c, stationId)\n"))
		if resource.RequestType != "" {
			f.WriteString(fmt.Sprintf("\t\tres, err := req.%s(context.Background(), &metadataApi.%s{\n", resource.Name, resource.RequestType))
			for _, fieldDef := range reqModel.Fields {
				switch fieldDef.Type {
				case "DateParam":
					f.WriteString(fmt.Sprintf("\t\t\t%s:  util.ParseDateParam(%sBeginDate, %sEndDate, %sRangeHours, %sRelative),\n", ToUpperCamelCase(fieldDef.Name), ToCamelCase(fieldDef.Name), ToCamelCase(fieldDef.Name), ToCamelCase(fieldDef.Name), ToCamelCase(fieldDef.Name)))
				case "IntervalParam":
					f.WriteString(fmt.Sprintf("\t\t\t%s:  util.ParseIntervalParam(%s),\n", ToUpperCamelCase(fieldDef.Name), ToCamelCase(fieldDef.Name)))
				case "VelocityTypeParam":
				default:
					f.WriteString(fmt.Sprintf("\t\t\t%s: %s,\n", fieldDef.Name, ToCamelCase(fieldDef.Name)))
				}
			}
			f.WriteString(fmt.Sprintf("\t\t})\n"))
		} else {
			f.WriteString(fmt.Sprintf("\t\tres, err := req.%s(context.Background())\n", resource.Name))
		}
		f.WriteString(fmt.Sprintf("\t\tif err != nil {\n"))
		f.WriteString(fmt.Sprintf("\t\t\tlog.Fatal(err)\n"))
		f.WriteString(fmt.Sprintf("\t\t}\n"))
		f.WriteString(fmt.Sprintf("\t\tjsonBytes, _ := json.MarshalIndent(res, \"\", \"  \")\n"))
		f.WriteString(fmt.Sprintf("\t\tfmt.Printf(\"%%s\\n\", jsonBytes)\n"))
		f.WriteString(fmt.Sprintf("\t},\n"))
		f.WriteString("}\n\n")

		// write init
		f.WriteString("func init() {\n")
		f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&stationId, \"station-id\", \"\", \"station id\")\n", ToUpperCamelCase(resource.ResourceID)))
		for _, fieldDef := range reqModel.Fields {
			fieldDesc := ""
			switch fieldDef.Type {
			case "DateParam":
				if hasMultDateParams {
					f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&%sBeginDate, \"%s-begin\", \"\", \"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToCamelCase(fieldDef.Name), ToKebabCase(fieldDef.Name), fieldDesc))
					f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&%sEndDate, \"%s-end\", \"\", \"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToCamelCase(fieldDef.Name), ToKebabCase(fieldDef.Name), fieldDesc))
					f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&%sRangeHours, \"%s-hours\", \"\", \"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToCamelCase(fieldDef.Name), ToKebabCase(fieldDef.Name), fieldDesc))
					f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&%sRelative, \"%s-relative\", \"\", \"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToCamelCase(fieldDef.Name), ToKebabCase(fieldDef.Name), fieldDesc))
				} else {
					f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&%sBeginDate, \"begin\", \"\", \"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToCamelCase(fieldDef.Name), fieldDesc))
					f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&%sEndDate, \"end\", \"\", \"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToCamelCase(fieldDef.Name), fieldDesc))
					f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&%sRangeHours, \"hours\", \"\", \"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToCamelCase(fieldDef.Name), fieldDesc))
					f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&%sRelative, \"relative\", \"\", \"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToCamelCase(fieldDef.Name), fieldDesc))
				}

			case "IntervalParam":
				f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&%s, \"%s\", string(metadataApi.%s), \"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToCamelCase(fieldDef.Name), ToKebabCase(fieldDef.Name), fieldDef.Default, fieldDesc))

			case "VelocityTypeParam":
			case "string":
				f.WriteString(fmt.Sprintf("\t%sCmd.Flags().StringVar(&%s, \"%s\", \"%s\", \"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToCamelCase(fieldDef.Name), ToKebabCase(fieldDef.Name), fieldDef.Default, fieldDesc))
			}
		}
		f.WriteString(fmt.Sprintf("\t%sCmd.MarkFlagRequired(\"station-id\")\n", ToUpperCamelCase(resource.ResourceID)))
		for _, fieldDef := range reqModel.Fields {
			if fieldDef.Required && fieldDef.Default == "" {
				switch fieldDef.Type {
				case "DateParam":
				case "IntervalParam":
				case "VelocityTypeParam":
				default:
					f.WriteString(fmt.Sprintf("\t%sCmd.MarkFlagRequired(\"%s\")\n", ToUpperCamelCase(resource.ResourceID), ToUpperCamelCase(fieldDef.Name)))
				}
			}
		}
		f.WriteString("}\n\n")
	}

	return nil
}
