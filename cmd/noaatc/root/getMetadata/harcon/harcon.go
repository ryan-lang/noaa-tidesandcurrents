package harcon

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"github.com/spf13/cobra"
	"github.com/ryan-lang/noaa-tidesandcurrents/client/metadataApi"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

var (
	stationId string
	units string
)

var HarconCmd = &cobra.Command{
	Use:   "harcon",
	Short: "Get harcon data",
	Long: `Get harcon data`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := metadataApi.NewClient(verbose, "github.com/ryan-lang/noaa-tidesandcurrents")
		req := metadataApi.NewStationRequest(c, stationId)
		res, err := req.HarmonicConstituents(context.Background(), &metadataApi.HarmonicConstituentsRequest{
			Units: units,
		})
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	HarconCmd.Flags().StringVar(&stationId, "station-id", "", "station id")
	HarconCmd.Flags().StringVar(&units, "units", "metric", "")
	HarconCmd.MarkFlagRequired("station-id")
}

