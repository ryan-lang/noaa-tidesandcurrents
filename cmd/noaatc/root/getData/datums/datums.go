package datums

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"github.com/spf13/cobra"
	"github.com/ryan-lang/noaa-tidesandcurrents/client/dataApi"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

var (
	stationId string
	units string
)

var DatumsCmd = &cobra.Command{
	Use:   "datums",
	Short: "Get datums data",
	Long: `Get datums data`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := dataApi.NewClient(verbose, "github.com/ryan-lang/noaa-tidesandcurrents")
		req := &dataApi.DatumsRequest{
			StationID: stationId,
			Units: units,
		}
		res, err := c.Datums(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	DatumsCmd.Flags().StringVar(&stationId, "station-id", "", "")
	DatumsCmd.Flags().StringVar(&units, "units", "Metric", "")
	DatumsCmd.MarkFlagRequired("StationId")
}

