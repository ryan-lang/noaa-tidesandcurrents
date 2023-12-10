package airPressure

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

var AirPressureCmd = &cobra.Command{
	Use:   "air-pressure",
	Short: "Get air_pressure data",
	Long: `Get air_pressure data`,
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
		res, err := c.AirPressure(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	AirPressureCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	AirPressureCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	AirPressureCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	AirPressureCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	AirPressureCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	AirPressureCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	AirPressureCmd.Flags().StringVar(&stationId, "station-id", "", "")
	AirPressureCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	AirPressureCmd.Flags().StringVar(&units, "units", "Metric", "")
	AirPressureCmd.MarkFlagRequired("StationId")
}

