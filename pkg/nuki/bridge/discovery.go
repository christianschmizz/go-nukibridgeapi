package bridge

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	discoveryUrl string = "https://api.nuki.io/discover/bridges"
)

type BridgeInfo struct {
	BridgeID    int       `json:"bridgeId,"`
	IP          string    `json:"ip,"`
	Port        int       `json:"port,"`
	DateUpdated time.Time `json:"dateUpdated,"`
}

type Discovery struct {
	Bridges   []BridgeInfo `json:"bridges"`
	ErrorCode int          `json:"errorCode,"`
}

func Discover() (*Discovery, error) {
	resp, err := http.Get(discoveryUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load discovery data from portal at %s", discoveryUrl)
	}
	defer resp.Body.Close()

	var discovery Discovery
	if err := json.NewDecoder(resp.Body).Decode(&discovery); err != nil {
		return nil, errors.Wrapf(err, "failed to decode discovery data")
	}

	return &discovery, nil
}
