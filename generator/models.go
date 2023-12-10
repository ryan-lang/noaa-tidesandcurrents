package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GenerateModels(path string, packageName string, models []ModelDefinition) error {

	// open file for writing; overrite if exists
	_ = os.MkdirAll(filepath.Dir(path), 0700)
	log.Printf("generating data api model file: %s", path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	// write package
	f.WriteString(fmt.Sprintf("package %s\n\n", packageName))

	// write gen warning
	f.WriteString("// THIS FILE IS GENERATED. DO NOT EDIT.\n\n")

	var requiredImports map[string]bool = map[string]bool{}
	var typeDefs, validators, unmarshallers []string

	// build the structs
	for _, modelDef := range models {
		typeDef, typeDefImports := modelTypeDefinition(modelDef)
		typeDefs = append(typeDefs, typeDef)
		for k, v := range typeDefImports {
			if v {
				requiredImports[k] = v
			}
		}
	}

	// build the validators
	for _, modelDef := range models {

		// skip the response models
		if modelDef.IsResponse {
			continue
		}

		validator, validatorImports := validatorFunc(modelDef)
		validators = append(validators, validator)
		for k, v := range validatorImports {
			if v {
				requiredImports[k] = v
			}
		}
	}

	// build the response unmarshallers
	for _, modelDef := range models {

		if !modelDef.IsResponse || modelDef.CustomUnmarshal {
			continue
		}

		unmarshalFunc, unmarshalImports := unmarshalFunc(modelDef)
		unmarshallers = append(unmarshallers, unmarshalFunc)
		for k, v := range unmarshalImports {
			if v {
				requiredImports[k] = v
			}
		}
	}

	// write imports
	f.WriteString("import (\n")
	for k := range requiredImports {
		f.WriteString(fmt.Sprintf("\t\"%s\"\n", k))
	}
	f.WriteString(")\n\n")

	// write the type definitions
	for _, typeDef := range typeDefs {
		f.WriteString(typeDef)
	}

	// write the validators
	for _, validator := range validators {
		f.WriteString(validator)
	}

	// write the unmarshallers
	for _, unmarshaller := range unmarshallers {
		f.WriteString(unmarshaller)
	}

	// gofmt the file
	err = goFmtFile(path)
	if err != nil {
		return err
	}

	return nil
}

func modelTypeDefinition(modelDef ModelDefinition) (string, map[string]bool) {
	var defString string
	var requiredImports map[string]bool = map[string]bool{}

	defString += fmt.Sprintf("type %s struct {\n", modelDef.Name)
	for _, fieldDef := range modelDef.Fields {

		defString += fmt.Sprintf("\t%s ", fieldDef.Name)
		if fieldDef.Required {
			defString += fieldDef.Type
		} else {
			defString += fmt.Sprintf("*%s ", fieldDef.Type)
		}
		if !modelDef.IsResponse {
			defString += fmt.Sprintf(" `url:\"%s\"`", fieldDef.GetUrlParam())
		}
		defString += "\n"
	}
	defString += "}\n\n"

	return defString, requiredImports
}

func validatorFunc(modelDef ModelDefinition) (string, map[string]bool) {
	var funcString string
	var requiredImports map[string]bool = map[string]bool{}

	funcString += fmt.Sprintf("func (m *%s) Validate() error {\n\n", modelDef.Name)
	for _, fieldDef := range modelDef.Fields {

		// special handling for certain types
		if fieldDef.Type == "DateParam" || fieldDef.Type == "IntervalParam" || fieldDef.Type == "VelocityTypeParam" {
			requiredImports["fmt"] = true
			if fieldDef.Default != "" {
				funcString += fmt.Sprintf("\tif m.%s == \"\" {\n", fieldDef.Name)
				funcString += fmt.Sprintf("\t\tm.%s = %s\n", fieldDef.Name, fieldDef.Default)
				funcString += "\t}\n\n"
			}

			if !fieldDef.Required || fieldDef.Type == "DateParam" {
				funcString += fmt.Sprintf("\tif m.%s != nil {\n", fieldDef.Name)
			}
			funcString += fmt.Sprintf("\tif err := m.%s.Validate(); err != nil {\n", fieldDef.Name)
			funcString += fmt.Sprintf("\t\treturn fmt.Errorf(\"%s parameter is invalid: %%w\", err)\n", strings.ToLower(fieldDef.Name))
			funcString += "\t}\n\n"
			if !fieldDef.Required && fieldDef.Type != "DateParam" {
				funcString += "\t}\n\n"
			} else if fieldDef.Type == "DateParam" {
				funcString += "\t} else {\n"
				funcString += fmt.Sprintf("\t\treturn fmt.Errorf(\"%s parameter is required\")\n", strings.ToLower(fieldDef.Name))
				funcString += "\t}\n\n"
			}

		} else if fieldDef.Default != "" {
			funcString += fmt.Sprintf("\tif m.%s == \"\" {\n", fieldDef.Name)
			funcString += fmt.Sprintf("\t\tm.%s = \"%s\"\n", fieldDef.Name, fieldDef.Default)
			funcString += "\t}\n\n"
		} else if fieldDef.Required {
			requiredImports["fmt"] = true
			funcString += fmt.Sprintf("\tif m.%s == \"\" {\n", fieldDef.Name)
			funcString += fmt.Sprintf("\t\treturn fmt.Errorf(\"%s is required\")\n", strings.ToLower(fieldDef.Name))
			funcString += "\t}\n\n"
		}
	}
	funcString += "\treturn nil\n"
	funcString += "}\n\n"

	return funcString, requiredImports
}

func unmarshalFunc(modelDef ModelDefinition) (string, map[string]bool) {

	var funcString string
	var requiredImports map[string]bool = map[string]bool{
		"encoding/json": true,
		"time":          false,
		"strconv":       false,
	}

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
			requiredImports["time"] = true
			requiredImports["fmt"] = true
			funcString += fmt.Sprintf("\t%sParsed, err := time.Parse(RESP_DATE_LAYOUT, %stmp.%s)\n", ToCamelCase(fieldDef.Name), deref, fieldDef.Name)
			funcString += "\tif err != nil {\n"
			funcString += fmt.Sprintf("\t\treturn fmt.Errorf(\"failed to parse %s: %%w\", err)\n", fieldDef.Name)
			funcString += "\t}\n\n"
		}
		// string -> float64
		if fieldDef.Type == "float64" && fieldDef.GetJsonType() == "string" {
			requiredImports["strconv"] = true
			requiredImports["fmt"] = true
			funcString += fmt.Sprintf("\t%sParsed, err := strconv.ParseFloat(%stmp.%s, 64)\n", ToCamelCase(fieldDef.Name), deref, fieldDef.Name)
			funcString += "\tif err != nil {\n"
			funcString += fmt.Sprintf("\t\treturn fmt.Errorf(\"failed to parse %s: %%w\", err)\n", fieldDef.Name)
			funcString += "\t}\n\n"
		}
		// string -> bool
		if fieldDef.Type == "bool" && fieldDef.GetJsonType() == "string" {
			requiredImports["strconv"] = true
			requiredImports["fmt"] = true
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
