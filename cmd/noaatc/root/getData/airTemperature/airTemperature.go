package airTemperature

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

var AirTemperatureCmd = &cobra.Command{
	Use:   "air-temperature",
	Short: "Get air_temperature data",
	Long: `Get air_temperature data`,
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
		res, err := c.AirTemperature(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	AirTemperatureCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	AirTemperatureCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	AirTemperatureCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	AirTemperatureCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	AirTemperatureCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	AirTemperatureCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	AirTemperatureCmd.Flags().StringVar(&stationId, "station-id", "", "")
	AirTemperatureCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	AirTemperatureCmd.Flags().StringVar(&units, "units", "Metric", "")
	AirTemperatureCmd.MarkFlagRequired("StationId")
}

