package dataApi_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ryan-lang/noaa-tidesandcurrents/client/dataApi"
)

func Example_tidePredictions() {

	// create a client
	c := dataApi.NewClient(false, "github.com/ryan-lang/noaa-tidesandcurrents")

	// create a request
	req := &dataApi.TidePredictionsRequest{
		StationID: "9447130",
		Date: &dataApi.DateParamBeginAndRange{
			BeginDate:  time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
			RangeHours: 1,
		},
		Interval: "10",
		Units:    "metric",
	}

	// make the request
	res, err := c.TidePredictions(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	// print the response
	for _, p := range res.Predictions {
		fmt.Printf("%f @ %s\n", p.Value, p.Time)
	}

	// Output:
	// 1.346000 @ 2023-04-10 00:00:00 +0000 UTC
	// 1.485000 @ 2023-04-10 00:10:00 +0000 UTC
	// 1.624000 @ 2023-04-10 00:20:00 +0000 UTC
	// 1.761000 @ 2023-04-10 00:30:00 +0000 UTC
	// 1.896000 @ 2023-04-10 00:40:00 +0000 UTC
	// 2.028000 @ 2023-04-10 00:50:00 +0000 UTC
	// 2.156000 @ 2023-04-10 01:00:00 +0000 UTC
}
