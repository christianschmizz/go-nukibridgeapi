package bridgeapi

import (
	"net/http"
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

// ScanResult represents a device found in reach of the bridge
type ScanResult struct {
	ID     int             `json:"nukiId"`
	Type   nuki.DeviceType `json:"deviceType"`
	Name   string          `json:"name"`
	Rssi   int             `json:"rssi"`
	Paired bool            `json:"paired"`
}

// NukiID assembles the ID from a result
func (r *ScanResult) NukiID() *nuki.ID {
	return &nuki.ID{
		DeviceID:   r.ID,
		DeviceType: r.Type,
	}
}

// InfoResponse represents the result of an info request
type InfoResponse struct {
	BridgeType BridgeType `json:"bridgeType"`
	IDs        struct {
		HardwareID int `json:"hardwareId"`
		ServerID   int `json:"serverId"`
	} `json:"ids"`
	Versions struct {
		FirmwareVersion     string `json:"firmwareVersion"`
		WifiFirmwareVersion string `json:"wifiFirmwareVersion"`
	} `json:"versions"`
	Uptime          int          `json:"uptime"`
	CurrentTime     time.Time    `json:"currentTime"`
	ServerConnected bool         `json:"serverConnected"`
	ScanResults     []ScanResult `json:"scanResults"`
}

// Info requests comprehensive information from the bridge
func (c *Connection) Info() (*InfoResponse, error) {
	resp, err := c.request("info", nil)
	if err != nil {
		return nil, err
	}

	if resp.Is(http.StatusUnauthorized) {
		return nil, ErrInvalidToken
	}

	var data InfoResponse
	if err := resp.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
