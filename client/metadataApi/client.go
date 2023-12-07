package metadataApi

import (
	"context"
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

func (c *Client) httpGet(ctx context.Context, urlPath string, params url.Values) ([]byte, error) {

	// build the url
	url := "https://api.tidesandcurrents.noaa.gov/mdapi/prod/webapi" + urlPath + "?" + params.Encode()

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
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	// read the response body into []byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	return body, nil
}
