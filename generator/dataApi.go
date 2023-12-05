package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GenerateDataApiClient(def DataApiDefinition) error {

	err := GenerateDataApiModels(def.Model)
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

func GenerateDataApiModels(models []ModelDefinition) error {

	// open file for writing; overrite if exists
	path := "./client/dataApi/model.gen.go"
	_ = os.MkdirAll(filepath.Dir(path), 0700)
	log.Printf("generating data api model file: %s", path)
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
	f.WriteString("\t\"fmt\"\n")
	f.WriteString("\t\"time\"\n")
	f.WriteString("\t\"encoding/json\"\n")
	f.WriteString("\t\"strconv\"\n")
	f.WriteString(")\n\n")

	// write the structs
	for _, modelDef := range models {
		f.WriteString(fmt.Sprintf("type %s struct {\n", modelDef.Name))
		for _, fieldDef := range modelDef.Fields {

			f.WriteString(fmt.Sprintf("\t%s ", fieldDef.Name))
			if fieldDef.Required {
				f.WriteString(fieldDef.Type)
			} else {
				f.WriteString(fmt.Sprintf("*%s ", fieldDef.Type))
			}
			if !modelDef.IsResponse {
				f.WriteString(fmt.Sprintf(" `url:\"%s\"`", fieldDef.GetUrlParam()))
			}
			f.WriteString("\n")
		}
		f.WriteString("}\n\n")
	}

	// write the validators
	for _, modelDef := range models {

		// skip the response models
		if modelDef.IsResponse {
			continue
		}

		f.WriteString(fmt.Sprintf("func (m *%s) Validate() error {\n\n", modelDef.Name))
		for _, fieldDef := range modelDef.Fields {

			// special handling for certain types
			if fieldDef.Type == "DateParam" || fieldDef.Type == "IntervalParam" || fieldDef.Type == "VelocityTypeParam" {
				if fieldDef.Default != "" {
					f.WriteString(fmt.Sprintf("\tif m.%s == \"\" {\n", fieldDef.Name))
					f.WriteString(fmt.Sprintf("\t\tm.%s = %s\n", fieldDef.Name, fieldDef.Default))
					f.WriteString("\t}\n\n")
				}

				if !fieldDef.Required {
					f.WriteString(fmt.Sprintf("\tif m.%s != nil {\n", fieldDef.Name))
				}
				f.WriteString(fmt.Sprintf("\tif err := m.%s.Validate(); err != nil {\n", fieldDef.Name))
				f.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"%s parameter is invalid: %%w\", err)\n", strings.ToLower(fieldDef.Name)))
				f.WriteString("\t}\n\n")
				if !fieldDef.Required {
					f.WriteString("\t}\n\n")
				}

			} else if fieldDef.Default != "" {
				f.WriteString(fmt.Sprintf("\tif m.%s == \"\" {\n", fieldDef.Name))
				f.WriteString(fmt.Sprintf("\t\tm.%s = \"%s\"\n", fieldDef.Name, fieldDef.Default))
				f.WriteString("\t}\n\n")
			} else if fieldDef.Required {
				f.WriteString(fmt.Sprintf("\tif m.%s == \"\" {\n", fieldDef.Name))
				f.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"%s is required\")\n", strings.ToLower(fieldDef.Name)))
				f.WriteString("\t}\n\n")
			}
		}
		f.WriteString("\treturn nil\n")
		f.WriteString("}\n\n")
	}

	// write the response unmarshallers
	for _, modelDef := range models {

		if !modelDef.IsResponse || modelDef.CustomUnmarshal {
			continue
		}

		unmarshalFunc, _ := unmarshalFunc(modelDef)
		f.WriteString(unmarshalFunc)
		f.WriteString("\n")
	}

	// gofmt the file
	err = goFmtFile(path)
	if err != nil {
		return err
	}

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

func unmarshalFunc(modelDef ModelDefinition) (string, []string) {

	var funcString string
	var requiredImports []string

	funcString += fmt.Sprintf("func (m *%s) UnmarshalJSON(b []byte) error {\n", modelDef.Name)
	funcString += "\tvar tmp struct {\n"
	for _, fieldDef := range modelDef.Fields {
		var pointerStr string

		if !fieldDef.Required {
			pointerStr = "*"
		}

		funcString += fmt.Sprintf("\t\t%s %s%s `json:\"%s\"`\n", fieldDef.Name, pointerStr, fieldDef.GetJsonType(), fieldDef.GetJsonParam())
	}
	funcString += "\t}\n"
	funcString += "\terr := json.Unmarshal(b, &tmp)\n"
	funcString += "\tif err != nil {\n"
	funcString += "\t\treturn err\n"
	funcString += "\t}\n\n"
	for _, fieldDef := range modelDef.Fields {

		var deref string
		if !fieldDef.Required {
			funcString += fmt.Sprintf("\tif tmp.%s != nil {\n", fieldDef.Name)
			deref = "*"
		}

		// do conversions from json type to go type ==
		// string -> *pointer (no need to do other conversions after this one)
		if !fieldDef.Required && fieldDef.GetJsonType() == "string" {
			funcString += fmt.Sprintf("\t\tif *tmp.%s == \"\" {\n", fieldDef.Name)
			funcString += fmt.Sprintf("\t\t\tm.%s = nil\n", fieldDef.Name)
			funcString += "\t\t} else {\n"
		}
		// string -> time.Time
		if fieldDef.Type == "time.Time" && fieldDef.GetJsonType() == "string" {
			funcString += fmt.Sprintf("\t%sParsed, err := time.Parse(RESP_DATE_LAYOUT, %stmp.%s)\n", ToCamelCase(fieldDef.Name), deref, fieldDef.Name)
			funcString += "\tif err != nil {\n"
			funcString += fmt.Sprintf("\t\treturn fmt.Errorf(\"failed to parse %s: %%w\", err)\n", fieldDef.Name)
			funcString += "\t}\n\n"
		}
		// string -> float64
		if fieldDef.Type == "float64" && fieldDef.GetJsonType() == "string" {
			funcString += fmt.Sprintf("\t%sParsed, err := strconv.ParseFloat(%stmp.%s, 64)\n", ToCamelCase(fieldDef.Name), deref, fieldDef.Name)
			funcString += "\tif err != nil {\n"
			funcString += fmt.Sprintf("\t\treturn fmt.Errorf(\"failed to parse %s: %%w\", err)\n", fieldDef.Name)
			funcString += "\t}\n\n"
		}
		// string -> bool
		if fieldDef.Type == "bool" && fieldDef.GetJsonType() == "string" {
			funcString += fmt.Sprintf("\t%sParsed, err := strconv.ParseBool(%stmp.%s)\n", ToCamelCase(fieldDef.Name), deref, fieldDef.Name)
			funcString += "\tif err != nil {\n"
			funcString += fmt.Sprintf("\t\treturn fmt.Errorf(\"failed to parse %s: %%w\", err)\n", fieldDef.Name)
			funcString += "\t}\n\n"
		}

		if !fieldDef.Required {

			funcString += fmt.Sprintf("\tm.%s = ", fieldDef.Name)
			if fieldDef.Type != fieldDef.GetJsonType() {
				funcString += fmt.Sprintf("&%sParsed\n", ToCamelCase(fieldDef.Name))
			} else {
				funcString += fmt.Sprintf("tmp.%s\n", fieldDef.Name)
			}

			if !fieldDef.Required && fieldDef.GetJsonType() == "string" {
				// close the if block
				funcString += "\t\t}\n\n"
			}

			funcString += "\t}\n\n"
		}
	}

	for _, fieldDef := range modelDef.Fields {
		if fieldDef.Required {
			funcString += fmt.Sprintf("\tm.%s = ", fieldDef.Name)
			if fieldDef.Type != fieldDef.GetJsonType() {
				funcString += fmt.Sprintf("%sParsed\n", ToCamelCase(fieldDef.Name))
			} else {
				funcString += fmt.Sprintf("tmp.%s\n", fieldDef.Name)
			}
		}
	}

	funcString += "\treturn nil\n"
	funcString += "}\n\n"

	return funcString, requiredImports
}
