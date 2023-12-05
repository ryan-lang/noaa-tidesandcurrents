package generator

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/iancoleman/strcase"
)

func GenerateAll(def *ClientDefinition) error {

	err := GenerateDataApiClient(def.DataAPI)
	if err != nil {
		return err
	}

	err = GenerateMetadataApiClient(def.MetadataAPI)
	if err != nil {
		return err
	}

	return nil
}

func ToCamelCase(s string) string {
	return strcase.ToLowerCamel(s)
}

func ToUpperCamelCase(s string) string {
	return strcase.ToCamel(s)
}

func ToSnakeCase(s string) string {
	return strcase.ToSnake(s)
}

func goFmtFile(filename string) error {
	cmd := exec.Command("gofmt", "-w", filename)

	// Optional: Capture the output (both stdout and stderr)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Run the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running gofmt: %v, output: %s", err, out.String())
	}

	return nil
}
