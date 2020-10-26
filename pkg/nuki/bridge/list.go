package bridge

import (
	"fmt"
	"time"
)

type DeviceInfo struct {
	ID             int    `json:"nukiId"`
	Type           int    `json:"deviceType"`
	Name           string `json:"name"`
	LastKnownState struct {
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
	} `json:"lastKnownState"`
}

type ListResponse []DeviceInfo

// Returns a list of all paired Nuki scan
func (c *connection) ListPairedDevices() (*ListResponse, error) {
	var devices ListResponse
	if err := c.get(c.hashedURL("list", nil), &devices); err != nil {
		return nil, fmt.Errorf("could not fetch list of paired scan: %w", err)
	}
	return &devices, nil
}
