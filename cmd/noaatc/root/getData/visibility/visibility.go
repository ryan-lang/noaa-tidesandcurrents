package visibility

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

var VisibilityCmd = &cobra.Command{
	Use:   "visibility",
	Short: "Get visibility data",
	Long: `Get visibility data`,
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
		res, err := c.Visibility(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	VisibilityCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	VisibilityCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	VisibilityCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	VisibilityCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	VisibilityCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	VisibilityCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	VisibilityCmd.Flags().StringVar(&stationId, "station-id", "", "")
	VisibilityCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	VisibilityCmd.Flags().StringVar(&units, "units", "Metric", "")
	VisibilityCmd.MarkFlagRequired("StationId")
}

