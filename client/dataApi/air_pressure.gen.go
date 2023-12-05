package dataApi

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

func (c *Client) AirPressure(ctx context.Context, req *StandardRequest) (*AirPressureResponse, error) {

	// validate the request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// build the params
	params, _ := query.Values(req)
	params.Add("product", "air_pressure")

	// make the request
	respBody, err := c.httpGet(ctx, params)
	if err != nil {
		return nil, err
	}

	// parse the response
	var resp AirPressureResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	return &resp, nil
}

