package dataApi

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

func (c *Client) Conductivity(ctx context.Context, req *StandardRequest) (*ConductivityResponse, error) {

	// validate the request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// build the params
	params, _ := query.Values(req)
	params.Add("product", "conductivity")

	// make the request
	respBody, err := c.httpGet(ctx, params)
	if err != nil {
		return nil, err
	}

	// parse the response
	var resp ConductivityResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	return &resp, nil
}

