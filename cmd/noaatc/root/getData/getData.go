package getData

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"github.com/spf13/cobra"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/waterLevel"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/hourlyHeight"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/highLow"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/dailyMean"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/monthlyMean"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/oneMinuteWaterLevel"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/predictions"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/datums"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/airGap"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/airTemperature"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/waterTemperature"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/wind"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/airPressure"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/conductivity"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/visibility"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/humidity"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/salinity"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/currents"
	"github.com/ryan-lang/noaa-tidesandcurrents/cmd/noaatc/root/getData/currentsPredictions"
)

var GetDataCmd = &cobra.Command{
	Use:   "getData",
	Short: "Get data from NOAA CO-OPS Data API",
	Long: `Get data from NOAA CO-OPS Data API`,
}

func init() {
	GetDataCmd.AddCommand(waterLevel.WaterLevelCmd)
	GetDataCmd.AddCommand(hourlyHeight.HourlyHeightCmd)
	GetDataCmd.AddCommand(highLow.HighLowCmd)
	GetDataCmd.AddCommand(dailyMean.DailyMeanCmd)
	GetDataCmd.AddCommand(monthlyMean.MonthlyMeanCmd)
	GetDataCmd.AddCommand(oneMinuteWaterLevel.OneMinuteWaterLevelCmd)
	GetDataCmd.AddCommand(predictions.PredictionsCmd)
	GetDataCmd.AddCommand(datums.DatumsCmd)
	GetDataCmd.AddCommand(airGap.AirGapCmd)
	GetDataCmd.AddCommand(airTemperature.AirTemperatureCmd)
	GetDataCmd.AddCommand(waterTemperature.WaterTemperatureCmd)
	GetDataCmd.AddCommand(wind.WindCmd)
	GetDataCmd.AddCommand(airPressure.AirPressureCmd)
	GetDataCmd.AddCommand(conductivity.ConductivityCmd)
	GetDataCmd.AddCommand(visibility.VisibilityCmd)
	GetDataCmd.AddCommand(humidity.HumidityCmd)
	GetDataCmd.AddCommand(salinity.SalinityCmd)
	GetDataCmd.AddCommand(currents.CurrentsCmd)
	GetDataCmd.AddCommand(currentsPredictions.CurrentsPredictionsCmd)
}

