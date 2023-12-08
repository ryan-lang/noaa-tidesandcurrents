package metadataApi

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"encoding/json"
)

type HarmonicConstituent struct {
	Amplitude   float64
	Description string
	Name        string
	Number      int64
	PhaseGMT    float64
	PhaseLocal  float64
	Speed       float64
}

type HarmonicConstituentsRequest struct {
	Units string `url:"units"`
}

type HarmonicConstituentsResponse struct {
	HarmonicConstituents []HarmonicConstituent
	Units                string
}

type TidePredictionOffsetsResponse struct {
	HeightAdjustedType   string
	HeightOffsetHighTide float64
	HeightOffsetLowTide  float64
	RefStationID         string
	TimeOffsetHighTide   float64
	TimeOffsetLowTide    float64
	Type                 string
}

func (m *HarmonicConstituentsRequest) Validate() error {

	if m.Units == "" {
		m.Units = "metric"
	}

	return nil
}

func (m *HarmonicConstituent) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Amplitude   float64 `json:"amplitude"`
		Description string  `json:"description"`
		Name        string  `json:"name"`
		Number      int64   `json:"number"`
		PhaseGMT    float64 `json:"phase_gmt"`
		PhaseLocal  float64 `json:"phase_local"`
		Speed       float64 `json:"speed"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.Amplitude = tmp.Amplitude
	m.Description = tmp.Description
	m.Name = tmp.Name
	m.Number = tmp.Number
	m.PhaseGMT = tmp.PhaseGMT
	m.PhaseLocal = tmp.PhaseLocal
	m.Speed = tmp.Speed
	return nil
}

func (m *HarmonicConstituentsResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		HarmonicConstituents []HarmonicConstituent `json:"HarmonicConstituents"`
		Units                string                `json:"units"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.HarmonicConstituents = tmp.HarmonicConstituents
	m.Units = tmp.Units
	return nil
}

func (m *TidePredictionOffsetsResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		HeightAdjustedType   string  `json:"heightAdjustedType"`
		HeightOffsetHighTide float64 `json:"heightOffsetHighTide"`
		HeightOffsetLowTide  float64 `json:"heightOffsetLowTide"`
		RefStationID         string  `json:"refStationId"`
		TimeOffsetHighTide   float64 `json:"timeOffsetHighTide"`
		TimeOffsetLowTide    float64 `json:"timeOffsetLowTide"`
		Type                 string  `json:"type"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	m.HeightAdjustedType = tmp.HeightAdjustedType
	m.HeightOffsetHighTide = tmp.HeightOffsetHighTide
	m.HeightOffsetLowTide = tmp.HeightOffsetLowTide
	m.RefStationID = tmp.RefStationID
	m.TimeOffsetHighTide = tmp.TimeOffsetHighTide
	m.TimeOffsetLowTide = tmp.TimeOffsetLowTide
	m.Type = tmp.Type
	return nil
}
