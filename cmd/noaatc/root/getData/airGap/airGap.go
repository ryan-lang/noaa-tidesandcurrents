package airGap

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

var AirGapCmd = &cobra.Command{
	Use:   "air-gap",
	Short: "Get air_gap data",
	Long: `Get air_gap data`,
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
		res, err := c.AirGap(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	AirGapCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	AirGapCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	AirGapCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	AirGapCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	AirGapCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	AirGapCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	AirGapCmd.Flags().StringVar(&stationId, "station-id", "", "")
	AirGapCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	AirGapCmd.Flags().StringVar(&units, "units", "Metric", "")
	AirGapCmd.MarkFlagRequired("StationId")
}

