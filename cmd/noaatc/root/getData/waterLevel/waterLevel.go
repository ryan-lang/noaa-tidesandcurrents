package waterLevel

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

var WaterLevelCmd = &cobra.Command{
	Use:   "water-level",
	Short: "Get water_level data",
	Long: `Get water_level data`,
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
		res, err := c.WaterLevel(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	WaterLevelCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	WaterLevelCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	WaterLevelCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	WaterLevelCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	WaterLevelCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	WaterLevelCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	WaterLevelCmd.Flags().StringVar(&stationId, "station-id", "", "")
	WaterLevelCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	WaterLevelCmd.Flags().StringVar(&units, "units", "Metric", "")
	WaterLevelCmd.MarkFlagRequired("StationId")
}

