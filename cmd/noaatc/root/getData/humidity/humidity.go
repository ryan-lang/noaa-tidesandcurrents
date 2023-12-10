package humidity

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

var HumidityCmd = &cobra.Command{
	Use:   "humidity",
	Short: "Get humidity data",
	Long: `Get humidity data`,
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
		res, err := c.Humidity(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	HumidityCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	HumidityCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	HumidityCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	HumidityCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	HumidityCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	HumidityCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	HumidityCmd.Flags().StringVar(&stationId, "station-id", "", "")
	HumidityCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	HumidityCmd.Flags().StringVar(&units, "units", "Metric", "")
	HumidityCmd.MarkFlagRequired("StationId")
}

