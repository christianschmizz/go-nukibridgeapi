package bridgeapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	discoveryURL string = "https://api.nuki.io/discover/bridges"
)

// BridgeInfo contains the basic information of a bridge device
type BridgeInfo struct {
	BridgeID    int       `json:"bridgeId,"`
	IP          string    `json:"ip,"`
	Port        int       `json:"port,"`
	DateUpdated time.Time `json:"dateUpdated,"`
}

// DiscoverResponse represents the result of a discovery request
type DiscoverResponse struct {
	Bridges   []BridgeInfo `json:"bridges"`
	ErrorCode int          `json:"errorCode,"`
}

// Discover requests a list of registered bridges from the public nuki API at the web
func Discover() (*DiscoverResponse, error) {
	resp, err := http.Get(discoveryURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load discovery data from portal at %s", discoveryURL)
	}
	defer resp.Body.Close()

	var response DiscoverResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.Wrapf(err, "failed to decode discovery data")
	}

	if response.ErrorCode != 0 {
		return nil, errors.New("unknown error occurred when trying to discover devices.")
	}

	return &response, nil
}
