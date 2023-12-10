package highLow

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

var HighLowCmd = &cobra.Command{
	Use:   "high-low",
	Short: "Get high_low data",
	Long: `Get high_low data`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := dataApi.NewClient(verbose, "github.com/ryan-lang/noaa-tidesandcurrents")
		req := &dataApi.StandardRequest{
			Date:  util.ParseDateParam(dateBeginDate, dateEndDate, dateRangeHours, dateRelative),
			Datum: datum,
			Interval:  util.ParseIntervalParam(interval),
			StationID: stationId,
			TimeZone: timeZone,
			Units: units,
		}
		res, err := c.HighLow(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	HighLowCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	HighLowCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	HighLowCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	HighLowCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	HighLowCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	HighLowCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	HighLowCmd.Flags().StringVar(&stationId, "station-id", "", "")
	HighLowCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	HighLowCmd.Flags().StringVar(&units, "units", "Metric", "")
	HighLowCmd.MarkFlagRequired("StationId")
}

