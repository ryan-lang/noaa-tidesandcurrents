package salinity

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

var SalinityCmd = &cobra.Command{
	Use:   "salinity",
	Short: "Get salinity data",
	Long: `Get salinity data`,
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
		res, err := c.Salinity(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	SalinityCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	SalinityCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	SalinityCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	SalinityCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	SalinityCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	SalinityCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	SalinityCmd.Flags().StringVar(&stationId, "station-id", "", "")
	SalinityCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	SalinityCmd.Flags().StringVar(&units, "units", "Metric", "")
	SalinityCmd.MarkFlagRequired("StationId")
}

