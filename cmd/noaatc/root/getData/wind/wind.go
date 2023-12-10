package wind

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

var WindCmd = &cobra.Command{
	Use:   "wind",
	Short: "Get wind data",
	Long: `Get wind data`,
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
		res, err := c.Wind(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	WindCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	WindCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	WindCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	WindCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	WindCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	WindCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	WindCmd.Flags().StringVar(&stationId, "station-id", "", "")
	WindCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	WindCmd.Flags().StringVar(&units, "units", "Metric", "")
	WindCmd.MarkFlagRequired("StationId")
}

