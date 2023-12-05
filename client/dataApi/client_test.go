package dataApi_test

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ryan-lang/noaa-tidesandcurrents/client/dataApi"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWaterLevel(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9447130",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.WaterLevel(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("WaterLevel response: %s\n", jsonBytes)

}

func TestHourlyHeight(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 7 * 6 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9447130",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.HourlyHeight(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("HourlyHeight response: %s\n", jsonBytes)

}

func TestHighLow(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 7 * 6 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9447130",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.HighLow(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("HighLow response: %s\n", jsonBytes)

}

func TestDailyMean(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.DailyMeanRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 7 * 6 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9087077",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.DailyMean(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("DailyMean response: %s\n", jsonBytes)

}

func TestMonthlyMean(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 7 * 6 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9447130",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.MonthlyMean(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("MonthlyMean response: %s\n", jsonBytes)

}

func TestWaterLevelHiRes(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9447130",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.WaterLevelHiRes(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("WaterLevelHiRes response: %s\n", jsonBytes)

}

func TestTidePredictions(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.TidePredictionsRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9447130",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.TidePredictions(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("TidePredictions response: %s\n", jsonBytes)

}

func TestDatums(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.DatumsRequest{
		StationID: "9447130",
		Units:     "",
	}
	res, err := c.Datums(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("Datums response: %s\n", jsonBytes)

}

func TestAirGap(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "8575432",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.AirGap(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("AirGap response: %s\n", jsonBytes)

}

func TestAirTemperature(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9445958",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.AirTemperature(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("AirTemperature response: %s\n", jsonBytes)

}

func TestWaterTemperature(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9444900",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.WaterTemperature(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("WaterTemperature response: %s\n", jsonBytes)

}

func TestWind(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9445958",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.Wind(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("Wind response: %s\n", jsonBytes)

}

func TestAirPressure(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9445958",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.AirPressure(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("AirPressure response: %s\n", jsonBytes)

}

func TestConductivity(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "8737048",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.Conductivity(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("Conductivity response: %s\n", jsonBytes)

}

func TestVisibility(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9414797",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.Visibility(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("Visibility response: %s\n", jsonBytes)

}

func TestHumidity(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9099064",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.Humidity(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("Humidity response: %s\n", jsonBytes)

}

func TestSalinity(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.StandardRequest{
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "9410170",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.Salinity(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("Salinity response: %s\n", jsonBytes)

}

func TestCurrents(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.CurrentsRequest{
		Bin: "",
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:     "",
		Interval:  "",
		StationID: "ks0101",
		TimeZone:  "",
		Units:     "",
	}
	res, err := c.Currents(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("Currents response: %s\n", jsonBytes)

}

func TestCurrentsPredictions(t *testing.T) {
	c := dataApi.NewClient(true, "test")
	ctx := context.Background()
	req := &dataApi.CurrentsPredictionsRequest{
		Bin: "",
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Now().Add(-24 * 3 * time.Hour),
			RangeHours: 24,
		},
		Datum:        "",
		Interval:     "",
		StationID:    "ks0101",
		TimeZone:     "",
		Units:        "",
		VelocityType: "",
	}
	res, err := c.CurrentsPredictions(ctx, req)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("CurrentsPredictions response: %s\n", jsonBytes)

}
