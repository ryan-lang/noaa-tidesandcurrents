package metadataApi

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"github.com/pkg/errors"
	"github.com/google/go-querystring/query"
)

func (c *StationRequest) HarmonicConstituents(ctx context.Context, req *HarmonicConstituentsRequest) (*HarmonicConstituentsResponse, error) {

	// check the fetched metadata to see if the resource is available
	if c.Metadata != nil {
		var isResourceAvailable bool
		for _, stationType := range c.Metadata.StationTypes() {
			if stationType == "STATION_TYPE_WATER_LEVEL" || stationType == "STATION_TYPE_CURRENT" || stationType == "STATION_TYPE_CURRENT_PREDICTION" {
				isResourceAvailable = true
				break
			}
		}
		if !isResourceAvailable {
			log.Printf("fetched metadata incidicates HarmonicConstituents is not available for station %s", c.StationID)
		}
	} else {
		if c.client.Verbose {
			log.Printf("availability of HarmonicConstituents for station %s is unknown. call FetchMetadata() first. trying anyway...", c.StationID)
		}
	}

	// validate the request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// build the params
	params, _ := query.Values(req)
	// make the request
	respBody, err := c.client.httpGet(ctx, fmt.Sprintf("/stations/%s/harcon.json", c.StationID), params)
	if err != nil {
		return nil, err
	}

	// parse the response
	var resp HarmonicConstituentsResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	return &resp, nil
}

func (c *StationsRequest) HarmonicConstituents(ctx context.Context) ([]*HarmonicConstituentsResponse, error) {

		// TODO: not yet implemented
		return nil, nil
}

