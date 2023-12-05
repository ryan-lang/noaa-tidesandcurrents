package metadataApi_test

// THIS FILE IS GENERATED. DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ryan-lang/noaa-tidesandcurrents/client/metadataApi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHarmonicConstituents(t *testing.T) {
	c := metadataApi.NewClient(true, "test")
	req := metadataApi.NewStationRequest(c, "9447130")
	ctx := context.Background()
	res, err := req.HarmonicConstituents(ctx)
	assert.NoError(t, err)
	jsonBytes, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("HarmonicConstituents response: %s\n", jsonBytes)

}
