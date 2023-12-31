// Provides a thin wrapper around the public NOAA Tides & Currents APIs.
// The two supported APIs are:
//
// Data API: https://api.tidesandcurrents.noaa.gov/api/prod/
//
// Metadata API: https://api.tidesandcurrents.noaa.gov/mdapi/prod/
package noaatidesandcurrents

import (
	"github.com/ryan-lang/noaa-tidesandcurrents/client/dataApi"
	"github.com/ryan-lang/noaa-tidesandcurrents/client/metadataApi"
)

type Client struct {
	Data     *dataApi.Client
	Metadata *metadataApi.Client
}

func NewClient(verbose bool, appName string) *Client {
	return &Client{
		Data:     dataApi.NewClient(verbose, appName),
		Metadata: metadataApi.NewClient(verbose, appName),
	}
}
