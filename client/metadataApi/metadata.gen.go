package metadataApi

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

func (c *StationRequest) Metadata(ctx context.Context) (*StationResponse, error) {

	// make the request
	respBody, err := c.client.httpGet(ctx, fmt.Sprintf("/stations/%s.json", c.StationID))
	if err != nil {
		return nil, err
	}

	// parse the response
	var resp StationResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	return &resp, nil
}

func (c *StationsRequest) Metadata(ctx context.Context) ([]*StationResponse, error) {

		// TODO: not yet implemented
		return nil, nil
}

