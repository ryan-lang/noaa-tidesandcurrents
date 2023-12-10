package dataApi

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type AirGapData struct {
	DataFlags *AirGapDataFlags
	Sigma     *float64
	Time      time.Time
	Value     *float64
}

type AirGapResponse struct {
	Data     []AirGapData
	Metadata Metadata
}

type AirPressureData struct {
	Flags AirPressureDataFlags
	Time  time.Time
	Value float64
}

type AirPressureResponse struct {
	Data     []AirPressureData
	Metadata Metadata
}

type AirTemperatureData struct {
	Flags AirTemperatureDataFlags
	Time  time.Time
	Value float64
}

type AirTemperatureResponse struct {
	Data     []AirTemperatureData
	Metadata Metadata
}

type ConductivityData struct {
	Flags ConductivityDataFlags
	Time  time.Time
	Value float64
}

type ConductivityResponse struct {
	Data     []ConductivityData
	Metadata Metadata
}

type CurrentPrediction struct {
	Bin           string
	Depth         float64
	Direction     float64
	MeanEbbDir    float64
	MeanFloodDir  float64
	Speed         float64
	Time          time.Time
	VelocityMajor float64
}

type CurrentsData struct {
	Bin       *string
	Direction float64
	Speed     float64
	Time      time.Time
}

type CurrentsPredictionsData struct {
	Predictions []CurrentPrediction
	Units       string
}

type CurrentsPredictionsRequest struct {
	Bin          string            `url:"bin"`
	Date         DateParam         `url:"date"`
	Datum        string            `url:"datum"`
	Interval     IntervalParam     `url:"interval"`
	StationID    string            `url:"station"`
	TimeZone     string            `url:"time_zone"`
	Units        string            `url:"units"`
	VelocityType VelocityTypeParam `url:"velocity_type"`
}

type CurrentsPredictionsResponse struct {
	CurrentPredictions CurrentsPredictionsData
}

type CurrentsRequest struct {
	Bin       string        `url:"bin"`
	Date      DateParam     `url:"date"`
	Datum     string        `url:"datum"`
	Interval  IntervalParam `url:"interval"`
	StationID string        `url:"station"`
	TimeZone  string        `url:"time_zone"`
	Units     string        `url:"units"`
}

type CurrentsResponse struct {
	Data     []CurrentsData
	Metadata Metadata
}

type DailyMeanData struct {
	Flags DailyMeanDataFlags
	Time  time.Time
	Value float64
}

type DailyMeanRequest struct {
	Date      DateParam     `url:"date"`
	Datum     string        `url:"datum"`
	Interval  IntervalParam `url:"interval"`
	StationID string        `url:"station"`
	TimeZone  string        `url:"time_zone"`
	Units     string        `url:"units"`
}

type DailyMeanResponse struct {
	Data     []DailyMeanData
	Metadata Metadata
}

type Datum struct {
	Name  string
	Value float64
}

type DatumsRequest struct {
	StationID string `url:"station"`
	Units     string `url:"units"`
}

type DatumsResponse struct {
	Datums []Datum
}

type HighLowData struct {
	Flags HighLowDataFlags
	Time  time.Time
	Type  string
	Value float64
}

type HighLowResponse struct {
	Data     []HighLowData
	Metadata Metadata
}

type HourlyHeightData struct {
	Flags HourlyHeightDataFlags
	Sigma float64
	Time  time.Time
	Value float64
}

type HourlyHeightResponse struct {
	Data     []HourlyHeightData
	Metadata Metadata
}

type HumidityData struct {
	Flags HumidityDataFlags
	Time  time.Time
	Value *float64
}

type HumidityResponse struct {
	Data     []HumidityData
	Metadata Metadata
}

type Metadata struct {
	Latitude    float64
	Longitude   float64
	StationID   string
	StationName string
}

type MonthlyMeanData struct {
	DHQ      float64
	DLQ      float64
	DTL      float64
	GT       float64
	HWI      float64
	Highest  float64
	Inferred bool
	LWI      float64
	Lowest   float64
	MHHW     float64
	MHW      float64
	MLLW     float64
	MLW      float64
	MN       float64
	MSL      float64
	MTL      float64
	Month    string
	Year     string
}

type MonthlyMeanResponse struct {
	Data     []MonthlyMeanData
	Metadata Metadata
}

type SalinityData struct {
	Salinity float64
	Time     time.Time
}

type SalinityResponse struct {
	Data     []SalinityData
	Metadata Metadata
}

type StandardRequest struct {
	Date      DateParam     `url:"date"`
	Datum     string        `url:"datum"`
	Interval  IntervalParam `url:"interval"`
	StationID string        `url:"station"`
	TimeZone  string        `url:"time_zone"`
	Units     string        `url:"units"`
}

type TidePrediction struct {
	Time  time.Time
	Type  *string
	Value float64
}

type TidePredictionsRequest struct {
	Date      DateParam     `url:"date"`
	Datum     string        `url:"datum"`
	Interval  IntervalParam `url:"interval"`
	StationID string        `url:"station"`
	TimeZone  string        `url:"time_zone"`
	Units     string        `url:"units"`
}

type TidePredictionsResponse struct {
	Predictions []TidePrediction
}

type VisibilityData struct {
	Flags VisibilityDataFlags
	Time  time.Time
	Value float64
}

type VisibilityResponse struct {
	Data     []VisibilityData
	Metadata Metadata
}

type WaterLevelData struct {
	DataFlags    WaterLevelDataFlags
	QualityLevel QualityLevel
	Sigma        *float64
	Time         time.Time
	Value        float64
}

type WaterLevelResponse struct {
	Data     []WaterLevelData
	Metadata Metadata
}

type WaterTemperatureData struct {
	Flags WaterTemperatureDataFlags
	Time  time.Time
	Value float64
}

type WaterTemperatureResponse struct {
	Data     []WaterTemperatureData
	Metadata Metadata
}

type WindData struct {
	Direction         float64
	DirectionCardinal string
	Flags             WindDataFlags
	Gust              float64
	Speed             float64
	Time              time.Time
}

type WindResponse struct {
	Data     []WindData
	Metadata Metadata
}

func (m *CurrentsPredictionsRequest) Validate() error {

	if m.Bin == "" {
		m.Bin = "1"
	}

	if m.Date != nil {
		if err := m.Date.Validate(); err != nil {
			return fmt.Errorf("date parameter is invalid: %w", err)
		}

	} else {
		return fmt.Errorf("date parameter is required")
	}

	if m.Datum == "" {
		m.Datum = "MLLW"
	}

	if m.Interval == "" {
		m.Interval = INTERVAL_PARAM_6M
	}

	if err := m.Interval.Validate(); err != nil {
		return fmt.Errorf("interval parameter is invalid: %w", err)
	}

	if m.StationID == "" {
		return fmt.Errorf("stationid is required")
	}

	if m.TimeZone == "" {
		m.TimeZone = "GMT"
	}

	if m.Units == "" {
		m.Units = "Metric"
	}

	if m.VelocityType == "" {
		m.VelocityType = VELOCITY_TYPE_DEFAULT
	}

	if err := m.VelocityType.Validate(); err != nil {
		return fmt.Errorf("velocitytype parameter is invalid: %w", err)
	}

	return nil
}

func (m *CurrentsRequest) Validate() error {

	if m.Bin == "" {
		m.Bin = "1"
	}

	if m.Date != nil {
		if err := m.Date.Validate(); err != nil {
			return fmt.Errorf("date parameter is invalid: %w", err)
		}

	} else {
		return fmt.Errorf("date parameter is required")
	}

	if m.Datum == "" {
		m.Datum = "MLLW"
	}

	if m.Interval == "" {
		m.Interval = INTERVAL_PARAM_6M
	}

	if err := m.Interval.Validate(); err != nil {
		return fmt.Errorf("interval parameter is invalid: %w", err)
	}

	if m.StationID == "" {
		return fmt.Errorf("stationid is required")
	}

	if m.TimeZone == "" {
		m.TimeZone = "GMT"
	}

	if m.Units == "" {
		m.Units = "Metric"
	}

	return nil
}

func (m *DailyMeanRequest) Validate() error {

	if m.Date != nil {
		if err := m.Date.Validate(); err != nil {
			return fmt.Errorf("date parameter is invalid: %w", err)
		}

	} else {
		return fmt.Errorf("date parameter is required")
	}

	if m.Datum == "" {
		m.Datum = "IGLD"
	}

	if m.Interval == "" {
		m.Interval = INTERVAL_PARAM_6M
	}

	if err := m.Interval.Validate(); err != nil {
		return fmt.Errorf("interval parameter is invalid: %w", err)
	}

	if m.StationID == "" {
		return fmt.Errorf("stationid is required")
	}

	if m.TimeZone == "" {
		m.TimeZone = "LST"
	}

	if m.Units == "" {
		m.Units = "Metric"
	}

	return nil
}

func (m *DatumsRequest) Validate() error {

	if m.StationID == "" {
		return fmt.Errorf("stationid is required")
	}

	if m.Units == "" {
		m.Units = "Metric"
	}

	return nil
}

func (m *StandardRequest) Validate() error {

	if m.Date != nil {
		if err := m.Date.Validate(); err != nil {
			return fmt.Errorf("date parameter is invalid: %w", err)
		}

	} else {
		return fmt.Errorf("date parameter is required")
	}

	if m.Datum == "" {
		m.Datum = "MLLW"
	}

	if m.Interval == "" {
		m.Interval = INTERVAL_PARAM_1M
	}

	if err := m.Interval.Validate(); err != nil {
		return fmt.Errorf("interval parameter is invalid: %w", err)
	}

	if m.StationID == "" {
		return fmt.Errorf("stationid is required")
	}

	if m.TimeZone == "" {
		m.TimeZone = "GMT"
	}

	if m.Units == "" {
		m.Units = "Metric"
	}

	return nil
}

func (m *TidePredictionsRequest) Validate() error {

	if m.Date != nil {
		if err := m.Date.Validate(); err != nil {
			return fmt.Errorf("date parameter is invalid: %w", err)
		}

	} else {
		return fmt.Errorf("date parameter is required")
	}

	if m.Datum == "" {
		m.Datum = "MLLW"
	}

	if m.Interval == "" {
		m.Interval = INTERVAL_PARAM_6M
	}

	if err := m.Interval.Validate(); err != nil {
		return fmt.Errorf("interval parameter is invalid: %w", err)
	}

	if m.StationID == "" {
		return fmt.Errorf("stationid is required")
	}

	if m.TimeZone == "" {
		m.TimeZone = "GMT"
	}

	if m.Units == "" {
		m.Units = "Metric"
	}

	return nil
}

func (m *AirGapData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		DataFlags *AirGapDataFlags `json:"f"`
		Sigma     *string          `json:"s"`
		Time      string           `json:"t"`
		Value     *string          `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	if tmp.DataFlags != nil {
		m.DataFlags = tmp.DataFlags
	}

	if tmp.Sigma != nil {
		if *tmp.Sigma == "" {
			m.Sigma = nil
		} else {
			sigmaParsed, err := strconv.ParseFloat(*tmp.Sigma, 64)
			if err != nil {
				return fmt.Errorf("failed to parse Sigma: %w", err)
			}

			m.Sigma = &sigmaParsed
		}

	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	if tmp.Value != nil {
		if *tmp.Value == "" {
			m.Value = nil
		} else {
			valueParsed, err := strconv.ParseFloat(*tmp.Value, 64)
			if err != nil {
				return fmt.Errorf("failed to parse Value: %w", err)
			}

			m.Value = &valueParsed
		}

	}

	m.Time = timeParsed
	return nil
}

func (m *AirGapResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []AirGapData `json:"data"`
		Metadata Metadata     `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *AirPressureData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Flags AirPressureDataFlags `json:"f"`
		Time  string               `json:"t"`
		Value string               `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	valueParsed, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	m.Flags = tmp.Flags
	m.Time = timeParsed
	m.Value = valueParsed
	return nil
}

func (m *AirPressureResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []AirPressureData `json:"data"`
		Metadata Metadata          `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *AirTemperatureData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Flags AirTemperatureDataFlags `json:"f"`
		Time  string                  `json:"t"`
		Value string                  `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	valueParsed, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	m.Flags = tmp.Flags
	m.Time = timeParsed
	m.Value = valueParsed
	return nil
}

func (m *AirTemperatureResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []AirTemperatureData `json:"data"`
		Metadata Metadata             `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *ConductivityData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Flags ConductivityDataFlags `json:"f"`
		Time  string                `json:"t"`
		Value string                `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	valueParsed, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	m.Flags = tmp.Flags
	m.Time = timeParsed
	m.Value = valueParsed
	return nil
}

func (m *ConductivityResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []ConductivityData `json:"data"`
		Metadata Metadata           `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *CurrentPrediction) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Bin           string  `json:"bin"`
		Depth         string  `json:"depth"`
		Direction     float64 `json:"direction"`
		MeanEbbDir    float64 `json:"meanEbbDir"`
		MeanFloodDir  float64 `json:"meanFloodDir"`
		Speed         float64 `json:"speed"`
		Time          string  `json:"Time"`
		VelocityMajor float64 `json:"Velocity_Major"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	depthParsed, err := strconv.ParseFloat(tmp.Depth, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Depth: %w", err)
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	m.Bin = tmp.Bin
	m.Depth = depthParsed
	m.Direction = tmp.Direction
	m.MeanEbbDir = tmp.MeanEbbDir
	m.MeanFloodDir = tmp.MeanFloodDir
	m.Speed = tmp.Speed
	m.Time = timeParsed
	m.VelocityMajor = tmp.VelocityMajor
	return nil
}

func (m *CurrentsData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Bin       *string `json:"b"`
		Direction string  `json:"d"`
		Speed     string  `json:"s"`
		Time      string  `json:"t"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	if tmp.Bin != nil {
		if *tmp.Bin == "" {
			m.Bin = nil
		} else {
			m.Bin = tmp.Bin
		}

	}

	directionParsed, err := strconv.ParseFloat(tmp.Direction, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Direction: %w", err)
	}

	speedParsed, err := strconv.ParseFloat(tmp.Speed, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Speed: %w", err)
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	m.Direction = directionParsed
	m.Speed = speedParsed
	m.Time = timeParsed
	return nil
}

func (m *CurrentsPredictionsData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Predictions []CurrentPrediction `json:"cp"`
		Units       string              `json:"units"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Predictions = tmp.Predictions
	m.Units = tmp.Units
	return nil
}

func (m *CurrentsPredictionsResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		CurrentPredictions CurrentsPredictionsData `json:"current_predictions"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.CurrentPredictions = tmp.CurrentPredictions
	return nil
}

func (m *CurrentsResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []CurrentsData `json:"data"`
		Metadata Metadata       `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *DailyMeanData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Flags DailyMeanDataFlags `json:"f"`
		Time  string             `json:"t"`
		Value string             `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	valueParsed, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	m.Flags = tmp.Flags
	m.Time = timeParsed
	m.Value = valueParsed
	return nil
}

func (m *DailyMeanResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []DailyMeanData `json:"data"`
		Metadata Metadata        `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *Datum) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Name  string `json:"n"`
		Value string `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	valueParsed, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	m.Name = tmp.Name
	m.Value = valueParsed
	return nil
}

func (m *DatumsResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Datums []Datum `json:"datums"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Datums = tmp.Datums
	return nil
}

func (m *HighLowData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Flags HighLowDataFlags `json:"f"`
		Time  string           `json:"t"`
		Type  string           `json:"ty"`
		Value string           `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	valueParsed, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	m.Flags = tmp.Flags
	m.Time = timeParsed
	m.Type = tmp.Type
	m.Value = valueParsed
	return nil
}

func (m *HighLowResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []HighLowData `json:"data"`
		Metadata Metadata      `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *HourlyHeightData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Flags HourlyHeightDataFlags `json:"f"`
		Sigma string                `json:"s"`
		Time  string                `json:"t"`
		Value string                `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	sigmaParsed, err := strconv.ParseFloat(tmp.Sigma, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Sigma: %w", err)
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	valueParsed, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	m.Flags = tmp.Flags
	m.Sigma = sigmaParsed
	m.Time = timeParsed
	m.Value = valueParsed
	return nil
}

func (m *HourlyHeightResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []HourlyHeightData `json:"data"`
		Metadata Metadata           `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *HumidityData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Flags HumidityDataFlags `json:"f"`
		Time  string            `json:"t"`
		Value *string           `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	if tmp.Value != nil {
		if *tmp.Value == "" {
			m.Value = nil
		} else {
			valueParsed, err := strconv.ParseFloat(*tmp.Value, 64)
			if err != nil {
				return fmt.Errorf("failed to parse Value: %w", err)
			}

			m.Value = &valueParsed
		}

	}

	m.Flags = tmp.Flags
	m.Time = timeParsed
	return nil
}

func (m *HumidityResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []HumidityData `json:"data"`
		Metadata Metadata       `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *Metadata) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Latitude    string `json:"lat"`
		Longitude   string `json:"lon"`
		StationID   string `json:"id"`
		StationName string `json:"name"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	latitudeParsed, err := strconv.ParseFloat(tmp.Latitude, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Latitude: %w", err)
	}

	longitudeParsed, err := strconv.ParseFloat(tmp.Longitude, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Longitude: %w", err)
	}

	m.Latitude = latitudeParsed
	m.Longitude = longitudeParsed
	m.StationID = tmp.StationID
	m.StationName = tmp.StationName
	return nil
}

func (m *MonthlyMeanData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		DHQ      string `json:"DHQ"`
		DLQ      string `json:"DLQ"`
		DTL      string `json:"DTL"`
		GT       string `json:"GT"`
		HWI      string `json:"HWI"`
		Highest  string `json:"highest"`
		Inferred string `json:"inferred"`
		LWI      string `json:"LWI"`
		Lowest   string `json:"lowest"`
		MHHW     string `json:"MHHW"`
		MHW      string `json:"MHW"`
		MLLW     string `json:"MLLW"`
		MLW      string `json:"MLW"`
		MN       string `json:"MN"`
		MSL      string `json:"MSL"`
		MTL      string `json:"MTL"`
		Month    string `json:"month"`
		Year     string `json:"year"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	dhqParsed, err := strconv.ParseFloat(tmp.DHQ, 64)
	if err != nil {
		return fmt.Errorf("failed to parse DHQ: %w", err)
	}

	dlqParsed, err := strconv.ParseFloat(tmp.DLQ, 64)
	if err != nil {
		return fmt.Errorf("failed to parse DLQ: %w", err)
	}

	dtlParsed, err := strconv.ParseFloat(tmp.DTL, 64)
	if err != nil {
		return fmt.Errorf("failed to parse DTL: %w", err)
	}

	gtParsed, err := strconv.ParseFloat(tmp.GT, 64)
	if err != nil {
		return fmt.Errorf("failed to parse GT: %w", err)
	}

	hwiParsed, err := strconv.ParseFloat(tmp.HWI, 64)
	if err != nil {
		return fmt.Errorf("failed to parse HWI: %w", err)
	}

	highestParsed, err := strconv.ParseFloat(tmp.Highest, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Highest: %w", err)
	}

	inferredParsed, err := strconv.ParseBool(tmp.Inferred)
	if err != nil {
		return fmt.Errorf("failed to parse Inferred: %w", err)
	}

	lwiParsed, err := strconv.ParseFloat(tmp.LWI, 64)
	if err != nil {
		return fmt.Errorf("failed to parse LWI: %w", err)
	}

	lowestParsed, err := strconv.ParseFloat(tmp.Lowest, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Lowest: %w", err)
	}

	mhhwParsed, err := strconv.ParseFloat(tmp.MHHW, 64)
	if err != nil {
		return fmt.Errorf("failed to parse MHHW: %w", err)
	}

	mhwParsed, err := strconv.ParseFloat(tmp.MHW, 64)
	if err != nil {
		return fmt.Errorf("failed to parse MHW: %w", err)
	}

	mllwParsed, err := strconv.ParseFloat(tmp.MLLW, 64)
	if err != nil {
		return fmt.Errorf("failed to parse MLLW: %w", err)
	}

	mlwParsed, err := strconv.ParseFloat(tmp.MLW, 64)
	if err != nil {
		return fmt.Errorf("failed to parse MLW: %w", err)
	}

	mnParsed, err := strconv.ParseFloat(tmp.MN, 64)
	if err != nil {
		return fmt.Errorf("failed to parse MN: %w", err)
	}

	mslParsed, err := strconv.ParseFloat(tmp.MSL, 64)
	if err != nil {
		return fmt.Errorf("failed to parse MSL: %w", err)
	}

	mtlParsed, err := strconv.ParseFloat(tmp.MTL, 64)
	if err != nil {
		return fmt.Errorf("failed to parse MTL: %w", err)
	}

	m.DHQ = dhqParsed
	m.DLQ = dlqParsed
	m.DTL = dtlParsed
	m.GT = gtParsed
	m.HWI = hwiParsed
	m.Highest = highestParsed
	m.Inferred = inferredParsed
	m.LWI = lwiParsed
	m.Lowest = lowestParsed
	m.MHHW = mhhwParsed
	m.MHW = mhwParsed
	m.MLLW = mllwParsed
	m.MLW = mlwParsed
	m.MN = mnParsed
	m.MSL = mslParsed
	m.MTL = mtlParsed
	m.Month = tmp.Month
	m.Year = tmp.Year
	return nil
}

func (m *MonthlyMeanResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []MonthlyMeanData `json:"data"`
		Metadata Metadata          `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *SalinityData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Salinity string `json:"s"`
		Time     string `json:"t"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	salinityParsed, err := strconv.ParseFloat(tmp.Salinity, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Salinity: %w", err)
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	m.Salinity = salinityParsed
	m.Time = timeParsed
	return nil
}

func (m *SalinityResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []SalinityData `json:"data"`
		Metadata Metadata       `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *TidePrediction) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Time  string  `json:"t"`
		Type  *string `json:"type"`
		Value string  `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	if tmp.Type != nil {
		if *tmp.Type == "" {
			m.Type = nil
		} else {
			m.Type = tmp.Type
		}

	}

	valueParsed, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	m.Time = timeParsed
	m.Value = valueParsed
	return nil
}

func (m *TidePredictionsResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Predictions []TidePrediction `json:"predictions"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Predictions = tmp.Predictions
	return nil
}

func (m *VisibilityData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Flags VisibilityDataFlags `json:"f"`
		Time  string              `json:"t"`
		Value string              `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	valueParsed, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	m.Flags = tmp.Flags
	m.Time = timeParsed
	m.Value = valueParsed
	return nil
}

func (m *VisibilityResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []VisibilityData `json:"data"`
		Metadata Metadata         `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *WaterLevelResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []WaterLevelData `json:"data"`
		Metadata Metadata         `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *WaterTemperatureData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Flags WaterTemperatureDataFlags `json:"f"`
		Time  string                    `json:"t"`
		Value string                    `json:"v"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	valueParsed, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	m.Flags = tmp.Flags
	m.Time = timeParsed
	m.Value = valueParsed
	return nil
}

func (m *WaterTemperatureResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []WaterTemperatureData `json:"data"`
		Metadata Metadata               `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}

func (m *WindData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Direction         string        `json:"d"`
		DirectionCardinal string        `json:"dr"`
		Flags             WindDataFlags `json:"f"`
		Gust              string        `json:"g"`
		Speed             string        `json:"s"`
		Time              string        `json:"t"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	directionParsed, err := strconv.ParseFloat(tmp.Direction, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Direction: %w", err)
	}

	gustParsed, err := strconv.ParseFloat(tmp.Gust, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Gust: %w", err)
	}

	speedParsed, err := strconv.ParseFloat(tmp.Speed, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Speed: %w", err)
	}

	timeParsed, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	m.Direction = directionParsed
	m.DirectionCardinal = tmp.DirectionCardinal
	m.Flags = tmp.Flags
	m.Gust = gustParsed
	m.Speed = speedParsed
	m.Time = timeParsed
	return nil
}

func (m *WindResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Data     []WindData `json:"data"`
		Metadata Metadata   `json:"metadata"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Data = tmp.Data
	m.Metadata = tmp.Metadata
	return nil
}
