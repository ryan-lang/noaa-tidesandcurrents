package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GenerateDataApiClient(def DataApiDefinition) error {

	modelsPath := "./client/dataApi/model.gen.go"
	err := GenerateModels(modelsPath, "dataApi", def.Model)
	if err != nil {
		return err
	}

	for _, productDef := range def.Products {
		err = GenerateDataApiProduct(productDef)
		if err != nil {
			return err
		}
	}

	err = GenerateDataApiTests(def)
	if err != nil {
		return err
	}

	return nil
}

func GenerateDataApiProduct(productDef ProductDefinition) error {

	// open file for writing; overrite if exists
	path := fmt.Sprintf("./client/dataApi/%s.gen.go", productDef.ProductID)
	_ = os.MkdirAll(filepath.Dir(path), 0700)
	log.Printf("generating data api file: %s", path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	// write package
	f.WriteString("package dataApi\n\n")

	// write gen warning
	f.WriteString("// THIS FILE IS GENERATED. DO NOT EDIT.\n\n")

	// write imports
	f.WriteString("import (\n")
	f.WriteString("\t\"context\"\n")
	f.WriteString("\t\"encoding/json\"\n")
	f.WriteString("\t\"github.com/google/go-querystring/query\"\n")
	f.WriteString("\t\"github.com/pkg/errors\"\n")
	f.WriteString(")\n\n")

	// write the method signature
	f.WriteString(fmt.Sprintf("func (c *Client) %s(ctx context.Context, req *%s) (*%s, error) {\n\n", productDef.Name, productDef.RequestType, productDef.ResponseType))

	// validate the request
	f.WriteString("\t// validate the request\n")
	f.WriteString("\tif err := req.Validate(); err != nil {\n")
	f.WriteString("\t\treturn nil, err\n")
	f.WriteString("\t}\n\n")

	// build the params
	f.WriteString("\t// build the params\n")
	f.WriteString("\tparams, _ := query.Values(req)\n")
	f.WriteString("\tparams.Add(\"product\", \"" + productDef.ProductID + "\")\n\n")

	// make the request
	f.WriteString("\t// make the request\n")
	f.WriteString("\trespBody, err := c.httpGet(ctx, params)\n")
	f.WriteString("\tif err != nil {\n")
	f.WriteString("\t\treturn nil, err\n")
	f.WriteString("\t}\n\n")

	// parse the response
	f.WriteString("\t// parse the response\n")
	f.WriteString("\tvar resp " + productDef.ResponseType + "\n")
	f.WriteString("\terr = json.Unmarshal(respBody, &resp)\n")
	f.WriteString("\tif err != nil {\n")
	f.WriteString("\t\treturn nil, errors.Wrap(err, \"failed to parse response\")\n")
	f.WriteString("\t}\n\n")

	// return the response
	f.WriteString("\treturn &resp, nil\n")
	f.WriteString("}\n\n")

	return nil
}

func GenerateDataApiTests(def DataApiDefinition) error {

	// open file for writing; overrite if exists
	path := "./client/dataApi/client_test.go"
	_ = os.MkdirAll(filepath.Dir(path), 0700)
	log.Printf("generating data api test file: %s", path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	// write package
	f.WriteString("package dataApi_test\n\n")

	// write gen warning
	f.WriteString("// THIS FILE IS GENERATED. DO NOT EDIT.\n\n")

	// write imports
	f.WriteString("import (\n")
	f.WriteString("\t\"context\"\n")
	f.WriteString("\t\"testing\"\n")
	f.WriteString("\t\"time\"\n")
	f.WriteString("\t\"encoding/json\"\n")
	f.WriteString("\t\"fmt\"\n")
	f.WriteString("\t\"github.com/stretchr/testify/assert\"\n")
	f.WriteString("\t\"github.com/ryan-lang/noaa-tidesandcurrents/client/dataApi\"\n")
	f.WriteString(")\n\n")

	// write the tests
	for _, productDef := range def.Products {
		f.WriteString(fmt.Sprintf("func Test%s(t *testing.T) {\n", productDef.Name))
		f.WriteString("\tc := dataApi.NewClient(true, \"test\")\n")
		f.WriteString("\tctx := context.Background()\n")
		f.WriteString(fmt.Sprintf("\treq := &dataApi.%s{\n", productDef.RequestType))
		requestTypeDef := def.GetModel(productDef.RequestType)
		for _, fieldDef := range requestTypeDef.Fields {

			if fieldDef.Required {
				f.WriteString(fmt.Sprintf("\t\t%s: ", fieldDef.Name))

				if fieldDef.Name == "StationID" {
					if productDef.TestStation == "" {
						productDef.TestStation = "9447130" // tide station Seattle
					}
					f.WriteString(fmt.Sprintf("\"%s\",\n", productDef.TestStation))
				} else if fieldDef.Type == "DateParam" {
					f.WriteString("&dataApi.DateParamBeginAndRange{\n")
					if productDef.TestHistoricalRange {
						f.WriteString("\t\t\tBeginDate:  time.Now().Add(-24 * 7 * 6 * time.Hour),\n")
					} else {
						f.WriteString("\t\t\tBeginDate:  time.Now().Add(-24 *  3 * time.Hour),\n")
					}
					f.WriteString("\t\t\tRangeHours: 24,\n")
					f.WriteString("\t\t},\n")
				} else {
					f.WriteString("\"\",\n")
				}
			}
		}
		f.WriteString("\t}\n")
		f.WriteString(fmt.Sprintf("\tres, err := c.%s(ctx, req)\n", productDef.Name))
		f.WriteString("\tassert.NoError(t, err)\n")

		// write the response to json and print it
		f.WriteString("jsonBytes, _ := json.MarshalIndent(res, \"\", \"  \")\n")
		f.WriteString(fmt.Sprintf("fmt.Printf(\"%s response: %%s\\n\", jsonBytes)\n\n", productDef.Name))

		f.WriteString("}\n\n")
	}

	// gofmt the file
	err = goFmtFile(path)
	if err != nil {
		return err
	}

	return nil
}
