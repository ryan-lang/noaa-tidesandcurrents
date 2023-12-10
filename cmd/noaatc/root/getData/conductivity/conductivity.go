package conductivity

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

var ConductivityCmd = &cobra.Command{
	Use:   "conductivity",
	Short: "Get conductivity data",
	Long: `Get conductivity data`,
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
		res, err := c.Conductivity(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	ConductivityCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	ConductivityCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	ConductivityCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	ConductivityCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	ConductivityCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	ConductivityCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_1M), "")
	ConductivityCmd.Flags().StringVar(&stationId, "station-id", "", "")
	ConductivityCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	ConductivityCmd.Flags().StringVar(&units, "units", "Metric", "")
	ConductivityCmd.MarkFlagRequired("StationId")
}

