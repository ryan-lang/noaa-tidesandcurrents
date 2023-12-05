package metadataApi

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"encoding/json"
	"fmt"
)

type ResourceRef struct {
	Self string `url:"self"`
}

type StationMetadata struct {
	Details    *ResourceRef
	GreatLakes *bool
	ShefCode   *string
	Tidal      *bool
}

type StationResponse struct {
	Count    int
	Stations []StationMetadata
	Units    *string
}

func (m *ResourceRef) Validate() error {

	if m.Self == "" {
		return fmt.Errorf("self is required")
	}

	return nil
}

func (m *StationMetadata) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Details    *ResourceRef `json:"details"`
		GreatLakes *bool        `json:"greatlakes"`
		ShefCode   *string      `json:"shefcode"`
		Tidal      *bool        `json:"tidal"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	if tmp.Details != nil {
		m.Details = tmp.Details
	}

	if tmp.GreatLakes != nil {
		m.GreatLakes = tmp.GreatLakes
	}

	if tmp.ShefCode != nil {
		if *tmp.ShefCode == "" {
			m.ShefCode = nil
		} else {
			m.ShefCode = tmp.ShefCode
		}

	}

	if tmp.Tidal != nil {
		m.Tidal = tmp.Tidal
	}

	return nil
}

func (m *StationResponse) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Count    int               `json:"count"`
		Stations []StationMetadata `json:"stations"`
		Units    *string           `json:"units"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	if tmp.Units != nil {
		if *tmp.Units == "" {
			m.Units = nil
		} else {
			m.Units = tmp.Units
		}

	}

	m.Count = tmp.Count
	m.Stations = tmp.Stations
	return nil
}
