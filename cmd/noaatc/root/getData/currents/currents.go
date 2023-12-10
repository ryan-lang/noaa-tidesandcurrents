package currents

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
	bin string
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

var CurrentsCmd = &cobra.Command{
	Use:   "currents",
	Short: "Get currents data",
	Long: `Get currents data`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		c := dataApi.NewClient(verbose, "github.com/ryan-lang/noaa-tidesandcurrents")
		req := &dataApi.CurrentsRequest{
			Bin: bin,
			Date:  util.ParseDateParam(dateBeginDate, dateEndDate, dateRangeHours, dateRelative),
			Datum: datum,
			Interval:  util.ParseIntervalParam(interval),
			StationID: stationId,
			TimeZone: timeZone,
			Units: units,
		}
		res, err := c.Currents(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		jsonBytes, _ := json.MarshalIndent(res, "", "  ")
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	CurrentsCmd.Flags().StringVar(&bin, "bin", "1", "")
	CurrentsCmd.Flags().StringVar(&dateBeginDate, "begin", "", "")
	CurrentsCmd.Flags().StringVar(&dateEndDate, "end", "", "")
	CurrentsCmd.Flags().StringVar(&dateRangeHours, "hours", "", "")
	CurrentsCmd.Flags().StringVar(&dateRelative, "relative", "", "")
	CurrentsCmd.Flags().StringVar(&datum, "datum", "MLLW", "")
	CurrentsCmd.Flags().StringVar(&interval, "interval", string(dataApi.INTERVAL_PARAM_6M), "")
	CurrentsCmd.Flags().StringVar(&stationId, "station-id", "", "")
	CurrentsCmd.Flags().StringVar(&timeZone, "time-zone", "GMT", "")
	CurrentsCmd.Flags().StringVar(&units, "units", "Metric", "")
	CurrentsCmd.MarkFlagRequired("StationId")
}

