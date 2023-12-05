package dataApi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Client struct {
	Verbose    bool
	AppName    string
	httpClient *http.Client
}

func NewClient(verbose bool, appName string) *Client {
	return &Client{
		Verbose:    verbose,
		AppName:    appName,
		httpClient: &http.Client{},
	}
}

func (c *Client) httpGet(ctx context.Context, params url.Values) ([]byte, error) {

	// add default params
	params.Add("format", "json")
	params.Add("application", c.AppName)

	// build the url
	url := "https://api.tidesandcurrents.noaa.gov/api/prod/datagetter?" + params.Encode()

	// build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}
	req = req.WithContext(ctx)

	if c.Verbose {
		log.Printf("making data api request: %s", url)
	}

	// execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute request")
	}
	defer resp.Body.Close()

	// check the response
	if resp.StatusCode != http.StatusOK {
		errResp, err := parseErrorResponse(resp.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse error response")
		}
		return nil, fmt.Errorf("bad status code: %d (%s)", resp.StatusCode, errResp.Error.Message)
	}

	// read the response body into []byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	// peek at the response body to see if it contains an error
	errResp, err := parseErrorResponse(io.NopCloser(bytes.NewReader(body)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse error response")
	}
	if errResp.Error.Message != "" {
		return nil, fmt.Errorf("error response: %s", errResp.Error.Message)
	}

	return body, nil
}

func parseErrorResponse(body io.Reader) (*ErrorResponse, error) {
	var errResp ErrorResponse
	err := json.NewDecoder(body).Decode(&errResp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse error response")
	}
	return &errResp, nil
}
