package waterTemperature

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

var WaterTemperatureCmd = &cobra.Command{
	Use:   "water-temperature",
	Short: "Get water_temperature data",
	Long: `Get water_temperature data`,
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
		res, err := c.WaterTemperature(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	WaterTemperatureCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	WaterTemperatureCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	WaterTemperatureCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	WaterTemperatureCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	WaterTemperatureCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	WaterTemperatureCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	WaterTemperatureCmd.Flags().StringVar(&stationId, "station-id", "", "")
	WaterTemperatureCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	WaterTemperatureCmd.Flags().StringVar(&units, "units", "Metric", "")
	WaterTemperatureCmd.MarkFlagRequired("StationId")
}

