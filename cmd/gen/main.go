package main

import (
	"flag"
	"log"

	"github.com/pkg/errors"
	"github.com/ryan-lang/noaa-tidesandcurrents/generator"
)

func main() {

	definitionPath := flag.String("definition", "./spec/noaa_api.yaml", "path to client definition")
	flag.Parse()

	def, err := generator.ParseClientDefinition(*definitionPath)
	if err != nil {
		panic(errors.Wrap(err, "failed to parse client definition"))
	}

	err = def.Validate()
	if err != nil {
		panic(errors.Wrap(err, "invalid client definition"))
	}

	log.Print("generating data api client")

	err = generator.GenerateDataApiClient(def.DataAPI)
	if err != nil {
		panic(errors.Wrap(err, "failed to generate data api client"))
	}
}
