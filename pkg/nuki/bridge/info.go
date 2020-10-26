package bridge

import (
	"fmt"
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

type ScanResult struct {
	ID     int             `json:"nukiId"`
	Type   nuki.DeviceType `json:"deviceType"`
	Name   string          `json:"name"`
	Rssi   int             `json:"rssi"`
	Paired bool            `json:"paired"`
}

func (r *ScanResult) NukiID() *nuki.NukiID {
	return &nuki.NukiID{
		DeviceID:   r.ID,
		DeviceType: r.Type,
	}
}

type Info struct {
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

func (c *connection) Info() (*Info, error) {
	var info Info
	if err := c.get(c.hashedURL("info", nil), &info); err != nil {
		return nil, fmt.Errorf("could not fetch bridge info: %w", err)
	}
	return &info, nil
}
