package metadataApi

type (
	StationRequest struct {
		StationID string
		Metadata  *StationMetadata
		client    *Client
	}
	StationsRequest struct {
		StationIDs []string
		Metadata   *StationMetadata
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
