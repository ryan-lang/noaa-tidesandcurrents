package tidepredoffsets

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
)

var TidepredoffsetsCmd = &cobra.Command{
	Use:   "tidepredoffsets",
	Short: "Get tidepredoffsets data",
	Long: `Get tidepredoffsets data`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := metadataApi.NewClient(verbose, "github.com/ryan-lang/noaa-tidesandcurrents")
		req := metadataApi.NewStationRequest(c, stationId)
		res, err := req.TidePredictionOffsets(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	TidepredoffsetsCmd.Flags().StringVar(&stationId, "station-id", "", "station id")
	TidepredoffsetsCmd.MarkFlagRequired("station-id")
}

