package metadataApi

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"github.com/pkg/errors"
)

func (c *StationRequest) TidePredictionOffsets(ctx context.Context) (*TidePredictionOffsetsResponse, error) {

	// check the fetched metadata to see if the resource is available
	if c.Metadata != nil {
		var isResourceAvailable bool
		for _, stationType := range c.Metadata.StationTypes() {
			if stationType == "STATION_TYPE_WATER_LEVEL" {
				isResourceAvailable = true
				break
			}
		}
		if !isResourceAvailable {
			log.Printf("fetched metadata incidicates TidePredictionOffsets is not available for station %s", c.StationID)
		}
	} else {
		if c.client.Verbose {
			log.Printf("availability of TidePredictionOffsets for station %s is unknown. call FetchMetadata() first. trying anyway...", c.StationID)
		}
	}

	// make the request
	respBody, err := c.client.httpGet(ctx, fmt.Sprintf("/stations/%s/tidepredoffsets.json", c.StationID), nil)
	if err != nil {
		return nil, err
	}

	// parse the response
	var resp TidePredictionOffsetsResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	return &resp, nil
}

func (c *StationsRequest) TidePredictionOffsets(ctx context.Context) ([]*TidePredictionOffsetsResponse, error) {

		// TODO: not yet implemented
		return nil, nil
}

