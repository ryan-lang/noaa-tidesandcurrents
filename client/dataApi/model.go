package dataApi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
)

type RelativeDateOpt string
type IntervalParam string
type QualityLevel string
type VelocityTypeParam string

const (
	RESP_DATE_LAYOUT = "2006-01-02 15:04"
	REQ_DATE_LAYOUT  = "20060102 15:04"

	RELATIVE_DATE_OPT_TODAY  RelativeDateOpt = "today"
	RELATIVE_DATE_OPT_LATEST RelativeDateOpt = "latest"
	RELATIVE_DATE_OPT_RECENT RelativeDateOpt = "recent"

	INTERVAL_PARAM_HILO   IntervalParam = "hilo"
	INTERVAL_PARAM_HOURLY IntervalParam = "h"
	INTERVAL_PARAM_1M     IntervalParam = "1"
	INTERVAL_PARAM_5M     IntervalParam = "5"
	INTERVAL_PARAM_6M     IntervalParam = "6"
	INTERVAL_PARAM_10M    IntervalParam = "10"
	INTERVAL_PARAM_15M    IntervalParam = "15"
	INTERVAL_PARAM_30M    IntervalParam = "30"
	INTERVAL_PARAM_60M    IntervalParam = "60"

	QUALITY_LEVEL_PRELIMINARY QualityLevel = "p"
	QUALITY_LEVEL_VERIFIED    QualityLevel = "v"

	VELOCITY_TYPE_DEFAULT   VelocityTypeParam = "default"
	VELOCITY_TYPE_SPEED_DIR VelocityTypeParam = "speed_dir"
)

type (
	// different ways that date & time can be specified
	DateParam interface {
		Validate() error
		query.Encoder
	}
	DateParamBeginAndEnd struct {
		BeginDate time.Time `url:"begin_date"`
		EndDate   time.Time `url:"end_date"`
	}
	DateParamBeginAndRange struct {
		BeginDate  time.Time `url:"begin_date"`
		RangeHours int32     `url:"range"`
	}
	DateParamEndAndRange struct {
		EndDate    time.Time `url:"end_date"`
		RangeHours int32     `url:"range"`
	}
	DateRelative struct { // NOTE: only available for preliminary water level data, meteorological data and predictions.
		Relative RelativeDateOpt `url:"date"`
	}
	DateRange struct { // Note only available for preliminary water level data, meteorological data
		RangeHours int32 `url:"range"`
	}

	ErrorResponse struct {
		Error ErrorDetail `json:"error"`
	}
	ErrorDetail struct {
		Message string `json:"message"`
	}
)

func (r *DateParamBeginAndEnd) Validate() error {
	if r.BeginDate.IsZero() {
		return fmt.Errorf("begin date is required")
	}

	if r.EndDate.IsZero() {
		return fmt.Errorf("end date is required")
	}

	if r.BeginDate.After(r.EndDate) {
		return fmt.Errorf("begin date must be before end date")
	}

	return nil
}

func (r *DateParamBeginAndRange) Validate() error {
	if r.BeginDate.IsZero() {
		return fmt.Errorf("begin date is required")
	}

	// TODO: max range?
	if r.RangeHours <= 0 {
		return fmt.Errorf("range must be greater than 0")
	}

	return nil
}

func (r *DateParamEndAndRange) Validate() error {
	if r.EndDate.IsZero() {
		return fmt.Errorf("end date is required")
	}

	if r.RangeHours <= 0 {
		return fmt.Errorf("range must be greater than 0")
	}

	return nil
}

func (r *DateRelative) Validate() error {
	if r.Relative == "" {
		return fmt.Errorf("relative date option is required")
	}

	if r.Relative != RELATIVE_DATE_OPT_TODAY &&
		r.Relative != RELATIVE_DATE_OPT_LATEST &&
		r.Relative != RELATIVE_DATE_OPT_RECENT {
		return fmt.Errorf("invalid relative date option: %s", r.Relative)
	}

	return nil
}

func (r *DateRange) Validate() error {
	if r.RangeHours <= 0 {
		return fmt.Errorf("range must be greater than 0")
	}

	return nil
}

func (r *DateParamBeginAndEnd) AddToParams(params url.Values) {
	params.Add("begin_date", r.BeginDate.Format(REQ_DATE_LAYOUT))
	params.Add("end_date", r.EndDate.Format(REQ_DATE_LAYOUT))
}

func (r *DateParamBeginAndRange) AddToParams(params url.Values) {
	params.Add("begin_date", r.BeginDate.Format(REQ_DATE_LAYOUT))
	params.Add("range", fmt.Sprintf("%d", r.RangeHours))
}

func (r *DateParamEndAndRange) AddToParams(params url.Values) {
	params.Add("end_date", r.EndDate.Format(REQ_DATE_LAYOUT))
	params.Add("range", fmt.Sprintf("%d", r.RangeHours))
}

func (r *DateRelative) AddToParams(params url.Values) {
	params.Add("date", string(r.Relative))
}

func (r *DateRange) AddToParams(params url.Values) {
	params.Add("range", fmt.Sprintf("%d", r.RangeHours))
}

func (r *DateParamBeginAndRange) EncodeValues(k string, v *url.Values) error {
	v.Add("begin_date", r.BeginDate.Format(REQ_DATE_LAYOUT))
	v.Add("range", fmt.Sprintf("%d", r.RangeHours))
	return nil
}

func (i IntervalParam) Validate() error {
	// TODO: make sure its one of our known intervals
	return nil
}

func (v VelocityTypeParam) Validate() error {
	// TODO: make sure its one of our known velocity types
	return nil
}

// custom unmarshal for WaterLevelData
func (m *WaterLevelData) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Time         string          `json:"t"`
		Value        string          `json:"v"`
		Sigma        *string         `json:"s"`
		DataFlags    json.RawMessage `json:"f"`
		QualityLevel QualityLevel    `json:"q"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	timeTime, err := time.Parse(RESP_DATE_LAYOUT, tmp.Time)
	if err != nil {
		return fmt.Errorf("failed to parse Time: %w", err)
	}

	valueFloat, err := strconv.ParseFloat(tmp.Value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse Value: %w", err)
	}

	if tmp.Sigma != nil {
		sigmaFloat, err := strconv.ParseFloat(*tmp.Sigma, 64)
		if err != nil {
			return fmt.Errorf("failed to parse Sigma: %w", err)
		}

		m.Sigma = &sigmaFloat
	}

	m.Time = timeTime
	m.Value = valueFloat
	m.QualityLevel = tmp.QualityLevel

	if m.QualityLevel == QUALITY_LEVEL_PRELIMINARY {
		var dataFlags PreliminaryWaterLevelDataFlags
		err = json.Unmarshal(tmp.DataFlags, &dataFlags)
		if err != nil {
			return fmt.Errorf("failed to parse DataFlags: %w", err)
		}
		m.DataFlags = dataFlags
	} else if m.QualityLevel == QUALITY_LEVEL_VERIFIED {
		var dataFlags VerifiedWaterLevelDataFlags
		err = json.Unmarshal(tmp.DataFlags, &dataFlags)
		if err != nil {
			return fmt.Errorf("failed to parse DataFlags: %w", err)
		}
		m.DataFlags = dataFlags
	}

	return nil
}
