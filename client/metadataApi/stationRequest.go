package metadataApi

type (
	StationRequest struct {
		StationID string
		client    *Client
	}
	StationsRequest struct {
		StationIDs []string
		client     *Client
	}
)

func NewStationRequest(client *Client, stationID string) *StationRequest {
	return &StationRequest{
		client:    client,
		StationID: stationID,
	}
}

func NewStationsRequest(client *Client, stationIDs []string) *StationsRequest {
	return &StationsRequest{
		client:     client,
		StationIDs: stationIDs,
	}
}
