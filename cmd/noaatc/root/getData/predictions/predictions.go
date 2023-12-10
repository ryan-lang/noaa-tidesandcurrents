package predictions

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"github.com/spf13/cobra"
	"github.com/ryan-lang/noaa-tidesandcurrents/client/dataApi"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/util"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

var (
	dateBeginDate string
	dateEndDate string
	dateRangeHours string
	dateRelative string
	datum string
	interval string
	stationId string
	timeZone string
	units string
)

var PredictionsCmd = &cobra.Command{
	Use:   "predictions",
	Short: "Get predictions data",
	Long: `Get predictions data`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := dataApi.NewClient(verbose, "github.com/ryan-lang/noaa-tidesandcurrents")
		req := &dataApi.TidePredictionsRequest{
			Date:  util.ParseDateParam(dateBeginDate, dateEndDate, dateRangeHours, dateRelative),
			Datum: datum,
			Interval:  util.ParseIntervalParam(interval),
			StationID: stationId,
			TimeZone: timeZone,
			Units: units,
		}
		res, err := c.TidePredictions(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	PredictionsCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	PredictionsCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	PredictionsCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	PredictionsCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	PredictionsCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	PredictionsCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_6M), "")
	PredictionsCmd.Flags().StringVar(&stationId, "station-id", "", "")
	PredictionsCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	PredictionsCmd.Flags().StringVar(&units, "units", "Metric", "")
	PredictionsCmd.MarkFlagRequired("StationId")
}

