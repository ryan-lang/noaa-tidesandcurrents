# NOAA Tides and Currents

A go library and CLI that provides an unofficial, thin wrapper around the NOAA Tides and Currents APIs.

The bulk of the code is generated from yaml files found in the `./spec` directory, making it quick and easy to keep up with changes to the NOAA APIs.

## CLI

```bash
go install github.com/ryan-lang/noaa-tidesandcurrents

noaatc getData predictions --station-id 9447130 --begin yesterday --hours 24 --interval 60 --units feet

noaatc getMetadata harcon --station-id 9447130
```

## Library

```go
// data api
import "github.com/ryan-lang/noaa-tidesandcurrents/client/dataApi"

client := dataApi.NewClient(true, "yourapplication")

req := &dataApi.TidePredictionsRequest{
    StationID: "9447130",
    Date: &dataApi.DateParamBeginAndRange{
        BeginDate:  time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
        RangeHours: 1,
    },
    Interval: "10",
    Units:    "metric",
}

res, err := client.TidePredictions(context.Background(), req)
if err != nil {
    log.Fatal(err)
}

```

```go
// metadata api
import "github.com/ryan-lang/noaa-tidesandcurrents/client/metadataApi"

client := metadataApi.NewClient(true, "yourapplication")

req := metadataApi.NewStationRequest(client, "9447130")

res, err := req.HarmonicConstituents(context.Background(), &metadataApi.HarmonicConstituentsRequest{
    Units: "metric",
})
if err != nil {
    log.Fatal(err)
}

```

### Supported endpoints
#### Data API

| NOAA Product ID | Method Name | CLI Command
| --- | --- | --- |
| water_level | WaterLevel | water-level |
| hourly_height | HourlyHeight | hourly-height |
| high_low | HighLow | high-low |
| daily_mean | DailyMean | daily-mean |
| monthly_mean | MonthlyMean | monthly-mean |
| one_minute_water_level | WaterLevelHiRes | water-level-hi-res |
| predictions | TidePredictions | tide-predictions |
| datums | Datums | datums |
| air_gap | AirGap | air-gap |
| air_temperature | AirTemperature | air-temperature |
| water_temperature | WaterTemperature | water-temperature |
| wind | Wind | wind |
| air_pressure | AirPressure | air-pressure |
| conductivity | Conductivity | conductivity |
| visibility | Visibility | visibility |
| humidity | Humidity | humidity |
| salinity | Salinity | salinity |
| currents | Currents | currents |
| currents_predictions | CurrentsPredictions | currents-predictions |


#### Metadata API

| Resource ID | Method Name |
| --- | --- |
| harcon | HarmonicConstituents |
| tidepredoffsets | TidePredictionOffsets |


## Contributing

```bash
git clone git@github.com:ryan-lang/noaa-tidesandcurrents.git
cd noaa-tidesandcurrents
go run ./cmd/gen/main.go
```

