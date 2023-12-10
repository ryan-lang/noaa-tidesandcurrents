package dailyMean

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

var DailyMeanCmd = &cobra.Command{
	Use:   "daily-mean",
	Short: "Get daily_mean data",
	Long: `Get daily_mean data`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := dataApi.NewClient(verbose, "github.com/ryan-lang/noaa-tidesandcurrents")
		req := &dataApi.DailyMeanRequest{
			Date:  util.ParseDateParam(dateBeginDate, dateEndDate, dateRangeHours, dateRelative),
			Datum: datum,
			Interval:  util.ParseIntervalParam(interval),
			StationID: stationId,
			TimeZone: timeZone,
			Units: units,
		}
		res, err := c.DailyMean(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	DailyMeanCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	DailyMeanCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	DailyMeanCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	DailyMeanCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	DailyMeanCmd.Flags().StringVar(&datum, "datum", "IGLD", "")
	DailyMeanCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_6M), "")
	DailyMeanCmd.Flags().StringVar(&stationId, "station-id", "", "")
	DailyMeanCmd.Flags().StringVar(&timeZone, "time-zone", "LST", "")
	DailyMeanCmd.Flags().StringVar(&units, "units", "Metric", "")
	DailyMeanCmd.MarkFlagRequired("StationId")
}

