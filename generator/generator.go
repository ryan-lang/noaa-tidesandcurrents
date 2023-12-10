package generator

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

	err = GenerateReadme(def)
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

func ToKebabCase(s string) string {
	return strings.ToLower(strcase.ToKebab(s))
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

func GenerateReadme(def *ClientDefinition) error {

	// open file for writing; overrite if exists
	path := "./README.md"
	_ = os.MkdirAll(filepath.Dir(path), 0700)
	log.Printf("generating data api file: %s", path)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	// write the intro
	f.WriteString("# NOAA Tides and Currents\n\n")
	f.WriteString("A go library and CLI that provides an unofficial, thin wrapper around the NOAA Tides and Currents APIs.\n\n")
	f.WriteString("The bulk of the code is generated from yaml files found in the `./spec` directory, making it quick and easy to keep up with changes to the NOAA APIs.\n\n")
	f.WriteString("## CLI\n\n")
	f.WriteString("```bash\n")
	f.WriteString("go install github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc@latest\n\n")
	f.WriteString("noaatc getData predictions --station-id 9447130 --begin yesterday --hours 24 --interval 60 --units english\n\n")
	f.WriteString("noaatc getMetadata harcon --station-id 9447130\n")
	f.WriteString("```\n\n")
	f.WriteString("## Library\n\n")
	f.WriteString("```go\n")
	f.WriteString("// data api\n")
	f.WriteString("import \"github.com/ryan-lang/noaa-tidesandcurrents/client/dataApi\"\n\n")
	f.WriteString("client := dataApi.NewClient(true, \"yourapplication\")\n\n")
	f.WriteString("req := &dataApi.TidePredictionsRequest{\n")
	f.WriteString("    StationID: \"9447130\",\n")
	f.WriteString("    Date: &dataApi.DateParamBeginAndRange{\n")
	f.WriteString("        BeginDate:  time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),\n")
	f.WriteString("        RangeHours: 1,\n")
	f.WriteString("    },\n")
	f.WriteString("    Interval: \"10\",\n")
	f.WriteString("    Units:    \"metric\",\n")
	f.WriteString("}\n\n")
	f.WriteString("res, err := client.TidePredictions(context.Background(), req)\n")
	f.WriteString("if err != nil {\n")
	f.WriteString("    log.Fatal(err)\n")
	f.WriteString("}\n\n")
	f.WriteString("```\n\n")
	f.WriteString("```go\n")
	f.WriteString("// metadata api\n")
	f.WriteString("import \"github.com/ryan-lang/noaa-tidesandcurrents/client/metadataApi\"\n\n")
	f.WriteString("client := metadataApi.NewClient(true, \"yourapplication\")\n\n")
	f.WriteString("req := metadataApi.NewStationRequest(client, \"9447130\")\n\n")
	f.WriteString("res, err := req.HarmonicConstituents(context.Background(), &metadataApi.HarmonicConstituentsRequest{\n")
	f.WriteString("    Units: \"metric\",\n")
	f.WriteString("})\n")
	f.WriteString("if err != nil {\n")
	f.WriteString("    log.Fatal(err)\n")
	f.WriteString("}\n\n")
	f.WriteString("```\n\n")
	f.WriteString("### Supported endpoints\n")
	f.WriteString("#### Data API\n\n")
	f.WriteString("| NOAA Product ID | Method Name | CLI Command\n")
	f.WriteString("| --- | --- | --- |\n")
	for _, p := range def.DataAPI.Products {
		f.WriteString(fmt.Sprintf("| %s | %s | %s |\n", p.ProductID, p.Name, ToKebabCase(p.Name)))
	}
	f.WriteString("\n\n")
	f.WriteString("#### Metadata API\n\n")
	f.WriteString("| Resource ID | Method Name |\n")
	f.WriteString("| --- | --- |\n")
	for _, r := range def.MetadataAPI.StationResources {
		f.WriteString(fmt.Sprintf("| %s | %s |\n", r.ResourceID, r.Name))
	}
	f.WriteString("\n\n")
	f.WriteString("## Contributing\n\n")
	f.WriteString("```bash\n")
	f.WriteString("git clone git@github.com:ryan-lang/noaa-tidesandcurrents.git\n")
	f.WriteString("cd noaa-tidesandcurrents\n")
	f.WriteString("go run ./cmd/gen/main.go\n")
	f.WriteString("```\n\n")

	return nil
}
