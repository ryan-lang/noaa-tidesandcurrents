package dataApi

import (
	"errors"
	"strconv"
	"strings"
)

type (
	PreliminaryWaterLevelDataFlags struct {
		Outliers                      int  // Count of number of 1 second samples that fall outside a 3-sigma band about the mean
		FlatToleranceExceeded         bool // A flag that when set to true indicates that the flat tolerance limit was exceeded
		RateOfChangeToleranceExceeded bool // A flag that when set to true indicates that the rate of change tolerance limit was exceeded
		LimitExceeded                 bool // A flag that when set to true indicates that either the maximum or minimum expected limit was exceeded
	}

	VerifiedWaterLevelDataFlags struct {
		Inferred                      bool //A flag that when set to true indicates that the water level value has been inferred
		FlatToleranceExceeded         bool // A flag that when set to true indicates that the flat tolerance limit was exceeded
		RateOfChangeToleranceExceeded bool // A flag that when set to true indicates that the rate of change tolerance limit was exceeded
		LimitExceeded                 bool // A flag that when set to true indicates that either the maximum or minimum expected limit was exceeded
	}

	WaterLevelDataFlags interface{}

	HourlyHeightDataFlags struct {
		Inferred      bool //A flag that when set to true indicates that the water level value has been inferred
		LimitExceeded bool // A flag that when set to true indicates that either the maximum or minimum expected limit was exceeded
	}

	HighLowDataFlags struct {
		Inferred      bool //A flag that when set to true indicates that the water level value has been inferred
		LimitExceeded bool // A flag that when set to true indicates that either the maximum or minimum expected limit was exceeded
	}

	DailyMeanDataFlags struct {
		Inferred      bool //A flag that when set to true indicates that the water level value has been inferred
		LimitExceeded bool // A flag that when set to true indicates that either the maximum or minimum expected limit was exceeded
	}

	AirGapDataFlags struct {
		Outliers                      int  // Count of number of 1 second samples that fall outside a 3-sigma band about the mean
		FlatToleranceExceeded         bool // A flag that when set to true indicates that the flat tolerance limit was exceeded
		RateOfChangeToleranceExceeded bool // A flag that when set to true indicates that the rate of change tolerance limit was exceeded
		LimitExceeded                 bool // A flag that when set to true indicates that either the maximum or minimum expected limit was exceeded
	}

	WindDataFlags struct {
		MaxExceeded                   bool // A flag that when set to true indicates that the maximum limit was exceeded
		RateOfChangeToleranceExceeded bool // A flag that when set to true indicates that the rate of change tolerance limit was exceeded
	}

	AirPressureDataFlags struct {
		MaxExceeded                   bool // A flag that when set to true indicates that the maximum limit was exceeded
		MinExceeded                   bool // A flag that when set to true indicates that the minimum limit was exceeded
		RateOfChangeToleranceExceeded bool // A flag that when set to true indicates that the rate of change tolerance limit was exceeded
	}

	AirTemperatureDataFlags struct {
		MaxExceeded                   bool // A flag that when set to true indicates that the maximum limit was exceeded
		MinExceeded                   bool // A flag that when set to true indicates that the minimum limit was exceeded
		RateOfChangeToleranceExceeded bool // A flag that when set to true indicates that the rate of change tolerance limit was exceeded
	}

	VisibilityDataFlags struct {
		MaxExceeded                   bool // A flag that when set to true indicates that the maximum limit was exceeded
		MinExceeded                   bool // A flag that when set to true indicates that the minimum limit was exceeded
		RateOfChangeToleranceExceeded bool // A flag that when set to true indicates that the rate of change tolerance limit was exceeded
	}

	HumidityDataFlags struct {
		MaxExceeded                   bool // A flag that when set to true indicates that the maximum limit was exceeded
		MinExceeded                   bool // A flag that when set to true indicates that the minimum limit was exceeded
		RateOfChangeToleranceExceeded bool // A flag that when set to true indicates that the rate of change tolerance limit was exceeded
	}

	WaterTemperatureDataFlags struct {
		MaxExceeded                   bool // A flag that when set to true indicates that the maximum limit was exceeded
		MinExceeded                   bool // A flag that when set to true indicates that the minimum limit was exceeded
		RateOfChangeToleranceExceeded bool // A flag that when set to true indicates that the rate of change tolerance limit was exceeded
	}

	ConductivityDataFlags struct {
		MaxExceeded                   bool // A flag that when set to true indicates that the maximum limit was exceeded
		MinExceeded                   bool // A flag that when set to true indicates that the minimum limit was exceeded
		RateOfChangeToleranceExceeded bool // A flag that when set to true indicates that the rate of change tolerance limit was exceeded
	}
)

func (d *PreliminaryWaterLevelDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")

	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 4 {
		return errors.New("invalid format: expected 4 comma-separated values")
	}

	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.Outliers, err = strconv.Atoi(parts[0]); err != nil {
		return errors.New("invalid Outliers value")
	}

	if d.FlatToleranceExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid FlatToleranceExceeded value")
	}

	if d.RateOfChangeToleranceExceeded, err = strconv.ParseBool(parts[2]); err != nil {
		return errors.New("invalid RateOfChangeToleranceExceeded value")
	}

	if d.LimitExceeded, err = strconv.ParseBool(parts[3]); err != nil {
		return errors.New("invalid LimitExceeded value")
	}

	return nil
}

func (d *VerifiedWaterLevelDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")

	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 4 {
		return errors.New("invalid format: expected 4 comma-separated values")
	}

	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.Inferred, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid Inferred value")
	}

	if d.FlatToleranceExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid FlatToleranceExceeded value")
	}

	if d.RateOfChangeToleranceExceeded, err = strconv.ParseBool(parts[2]); err != nil {
		return errors.New("invalid RateOfChangeToleranceExceeded value")
	}

	if d.LimitExceeded, err = strconv.ParseBool(parts[3]); err != nil {
		return errors.New("invalid LimitExceeded value")
	}

	return nil
}

func (d *HourlyHeightDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")

	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 2 {
		return errors.New("invalid format: expected 2 comma-separated values")
	}

	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.Inferred, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid Inferred value")
	}

	if d.LimitExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid LimitExceeded value")
	}

	return nil
}

func (d *HighLowDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")

	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 2 {
		return errors.New("invalid format: expected 2 comma-separated values")
	}

	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.Inferred, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid Inferred value")
	}

	if d.LimitExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid LimitExceeded value")
	}

	return nil
}

func (d *DailyMeanDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")

	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 2 {
		return errors.New("invalid format: expected 2 comma-separated values")
	}

	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.Inferred, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid Inferred value")
	}

	if d.LimitExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid LimitExceeded value")
	}

	return nil
}

func (d *AirGapDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")

	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 4 {
		return errors.New("invalid format: expected 4 comma-separated values")
	}

	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.Outliers, err = strconv.Atoi(parts[0]); err != nil {
		return errors.New("invalid Outliers value")
	}

	if d.FlatToleranceExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid FlatToleranceExceeded value")
	}

	if d.RateOfChangeToleranceExceeded, err = strconv.ParseBool(parts[2]); err != nil {
		return errors.New("invalid RateOfChangeToleranceExceeded value")
	}

	if d.LimitExceeded, err = strconv.ParseBool(parts[3]); err != nil {
		return errors.New("invalid LimitExceeded value")
	}

	return nil
}

func (d *WindDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")

	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 2 {
		return errors.New("invalid format: expected 2 comma-separated values")
	}

	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.MaxExceeded, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid MaxExceeded value")
	}

	if d.RateOfChangeToleranceExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid RateOfChangeToleranceExceeded value")
	}

	return nil
}

func (d *AirPressureDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")

	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 3 {
		return errors.New("invalid format: expected 3 comma-separated values")
	}

	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.MaxExceeded, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid MaxExceeded value")
	}

	if d.MinExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid MinExceeded value")
	}

	if d.RateOfChangeToleranceExceeded, err = strconv.ParseBool(parts[2]); err != nil {
		return errors.New("invalid RateOfChangeToleranceExceeded value")
	}

	return nil
}

func (d *AirTemperatureDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")

	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 3 {
		return errors.New("invalid format: expected 3 comma-separated values")
	}

	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.MaxExceeded, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid MaxExceeded value")
	}

	if d.MinExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid MinExceeded value")
	}

	if d.RateOfChangeToleranceExceeded, err = strconv.ParseBool(parts[2]); err != nil {
		return errors.New("invalid RateOfChangeToleranceExceeded value")
	}

	return nil
}

func (d *VisibilityDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")
	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 3 {
		return errors.New("invalid format: expected 3 comma-separated values")
	}
	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.MaxExceeded, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid MaxExceeded value")
	}
	if d.MinExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid MinExceeded value")
	}
	if d.RateOfChangeToleranceExceeded, err = strconv.ParseBool(parts[2]); err != nil {
		return errors.New("invalid RateOfChangeToleranceExceeded value")
	}
	return nil
}

func (d *HumidityDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")
	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 3 {
		return errors.New("invalid format: expected 3 comma-separated values")
	}
	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.MaxExceeded, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid MaxExceeded value")
	}
	if d.MinExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid MinExceeded value")
	}
	if d.RateOfChangeToleranceExceeded, err = strconv.ParseBool(parts[2]); err != nil {
		return errors.New("invalid RateOfChangeToleranceExceeded value")
	}
	return nil
}

func (d *WaterTemperatureDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")
	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 3 {
		return errors.New("invalid format: expected 3 comma-separated values")
	}
	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.MaxExceeded, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid MaxExceeded value")
	}
	if d.MinExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid MinExceeded value")
	}
	if d.RateOfChangeToleranceExceeded, err = strconv.ParseBool(parts[2]); err != nil {
		return errors.New("invalid RateOfChangeToleranceExceeded value")
	}
	return nil
}

func (d *ConductivityDataFlags) UnmarshalJSON(data []byte) error {
	// Convert the data to a string and trim the leading and trailing quotes.
	strData := strings.Trim(string(data), "\"")
	// Split the string on commas.
	parts := strings.Split(strData, ",")
	if len(parts) != 3 {
		return errors.New("invalid format: expected 3 comma-separated values")
	}
	// Convert each part to the appropriate type and assign to the struct fields.
	var err error
	if d.MaxExceeded, err = strconv.ParseBool(parts[0]); err != nil {
		return errors.New("invalid MaxExceeded value")
	}
	if d.MinExceeded, err = strconv.ParseBool(parts[1]); err != nil {
		return errors.New("invalid MinExceeded value")
	}
	if d.RateOfChangeToleranceExceeded, err = strconv.ParseBool(parts[2]); err != nil {
		return errors.New("invalid RateOfChangeToleranceExceeded value")
	}
	return nil
}
