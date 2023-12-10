package hourlyHeight

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

var HourlyHeightCmd = &cobra.Command{
	Use:   "hourly-height",
	Short: "Get hourly_height data",
	Long: `Get hourly_height data`,
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
		res, err := c.HourlyHeight(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	HourlyHeightCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	HourlyHeightCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	HourlyHeightCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	HourlyHeightCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	HourlyHeightCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	HourlyHeightCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	HourlyHeightCmd.Flags().StringVar(&stationId, "station-id", "", "")
	HourlyHeightCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	HourlyHeightCmd.Flags().StringVar(&units, "units", "Metric", "")
	HourlyHeightCmd.MarkFlagRequired("StationId")
}

