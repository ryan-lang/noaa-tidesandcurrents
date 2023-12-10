package oneMinuteWaterLevel

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

var OneMinuteWaterLevelCmd = &cobra.Command{
	Use:   "one-minute-water-level",
	Short: "Get one_minute_water_level data",
	Long: `Get one_minute_water_level data`,
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
		res, err := c.WaterLevelHiRes(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	OneMinuteWaterLevelCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	OneMinuteWaterLevelCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	OneMinuteWaterLevelCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	OneMinuteWaterLevelCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	OneMinuteWaterLevelCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	OneMinuteWaterLevelCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	OneMinuteWaterLevelCmd.Flags().StringVar(&stationId, "station-id", "", "")
	OneMinuteWaterLevelCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	OneMinuteWaterLevelCmd.Flags().StringVar(&units, "units", "Metric", "")
	OneMinuteWaterLevelCmd.MarkFlagRequired("StationId")
}

