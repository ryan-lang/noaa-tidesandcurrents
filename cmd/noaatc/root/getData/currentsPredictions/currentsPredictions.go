package currentsPredictions

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
	bin string
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

var CurrentsPredictionsCmd = &cobra.Command{
	Use:   "currents-predictions",
	Short: "Get currents_predictions data",
	Long: `Get currents_predictions data`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := dataApi.NewClient(verbose, "github.com/ryan-lang/noaa-tidesandcurrents")
		req := &dataApi.CurrentsPredictionsRequest{
			Bin: bin,
			Date:  util.ParseDateParam(dateBeginDate, dateEndDate, dateRangeHours, dateRelative),
			Datum: datum,
			Interval:  util.ParseIntervalParam(interval),
			StationID: stationId,
			TimeZone: timeZone,
			Units: units,
		}
		res, err := c.CurrentsPredictions(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	CurrentsPredictionsCmd.Flags().StringVar(&bin, "bin", "1", "")
	CurrentsPredictionsCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	CurrentsPredictionsCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	CurrentsPredictionsCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	CurrentsPredictionsCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	CurrentsPredictionsCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	CurrentsPredictionsCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_6M), "")
	CurrentsPredictionsCmd.Flags().StringVar(&stationId, "station-id", "", "")
	CurrentsPredictionsCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	CurrentsPredictionsCmd.Flags().StringVar(&units, "units", "Metric", "")
	CurrentsPredictionsCmd.MarkFlagRequired("StationId")
}

