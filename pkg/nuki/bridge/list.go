package bridge

import (
	"fmt"
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

type LastKnownState struct {
	Mode                  int       `json:"mode"`
	State                 int       `json:"state"`
	StateName             string    `json:"stateName"`
	BatteryCritical       bool      `json:"batteryCritical"`
	BatteryCharging       bool      `json:"batteryCharging"`
	BatteryChargeState    int       `json:"batteryChargeState"`
	KeypadBatteryCritical bool      `json:"keypadBatteryCritical,omitempty"`
	DoorsensorState       int       `json:"doorsensorState,omitempty"`
	DoorsensorStateName   string    `json:"doorsensorStateName,omitempty"`
	RingactionState       bool      `json:"ringactionState,omitempty"`
	RingactionTimestamp   time.Time `json:"ringactionTimestamp,omitempty"`
	Timestamp             time.Time `json:"timestamp"`
}
type DeviceInfo struct {
	ID             int             `json:"nukiId"`
	Type           nuki.DeviceType `json:"deviceType"`
	Name           string          `json:"name"`
	LastKnownState LastKnownState  `json:"lastKnownState"`
}

type ListPairedDevicesResponse []DeviceInfo

// Returns a list of all paired Nuki scan
func (c *Connection) ListPairedDevices() (ListPairedDevicesResponse, error) {
	var response ListPairedDevicesResponse
	if err := c.get(c.hashedURL("list", nil), &response); err != nil {
		return nil, fmt.Errorf("could not fetch list of paired scan: %w", err)
	}
	return response, nil
}
