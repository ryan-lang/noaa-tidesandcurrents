package monthlyMean

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

var MonthlyMeanCmd = &cobra.Command{
	Use:   "monthly-mean",
	Short: "Get monthly_mean data",
	Long: `Get monthly_mean data`,
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
		res, err := c.MonthlyMean(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	MonthlyMeanCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	MonthlyMeanCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	MonthlyMeanCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	MonthlyMeanCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	MonthlyMeanCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	MonthlyMeanCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	MonthlyMeanCmd.Flags().StringVar(&stationId, "station-id", "", "")
	MonthlyMeanCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	MonthlyMeanCmd.Flags().StringVar(&units, "units", "Metric", "")
	MonthlyMeanCmd.MarkFlagRequired("StationId")
}

