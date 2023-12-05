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

type HarmonicConstituentsResponse struct {
	HarmonicConstituents *[]HarmonicConstituent
	Units                string
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
		HarmonicConstituents *[]HarmonicConstituent `json:"HarmonicConstituents"`
		Units                string                 `json:"units"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	if tmp.HarmonicConstituents != nil {
		m.HarmonicConstituents = tmp.HarmonicConstituents
	}

	m.Units = tmp.Units
	return nil
}
