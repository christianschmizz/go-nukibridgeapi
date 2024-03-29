package bridgeapi

import (
	"net/http"
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

// LastKnownState describes the last known state of a device
type LastKnownState struct {
	Mode                  nuki.LockMode        `json:"mode"`
	State                 nuki.LockState       `json:"state"`
	StateName             string               `json:"stateName"`
	BatteryCritical       bool                 `json:"batteryCritical"`
	BatteryCharging       bool                 `json:"batteryCharging"`
	BatteryChargeState    uint8                `json:"batteryChargeState"`
	KeypadBatteryCritical bool                 `json:"keypadBatteryCritical,omitempty"`
	DoorsensorState       nuki.DoorsensorState `json:"doorsensorState,omitempty"`
	DoorsensorStateName   string               `json:"doorsensorStateName,omitempty"`
	RingactionState       bool                 `json:"ringactionState,omitempty"`
	RingactionTimestamp   time.Time            `json:"ringactionTimestamp,omitempty"`
	Timestamp             time.Time            `json:"timestamp"`
}

// DeviceInfo describes some basic information of a device
type DeviceInfo struct {
	ID              int             `json:"nukiId"`
	Type            nuki.DeviceType `json:"deviceType"`
	Name            string          `json:"name"`
	FirmwareVersion string          `json:"firmwareVersion"`
	LastKnownState  LastKnownState  `json:"lastKnownState"`
}

// ListPairedDevicesResponse represents the results of querying the paired devices
type ListPairedDevicesResponse []DeviceInfo

// ListPairedDevices retrieves a list of all devices paired with the bridge
func (c *Connection) ListPairedDevices() (ListPairedDevicesResponse, error) {
	resp, err := c.request("list", nil)
	if err != nil {
		return nil, err
	}

	if resp.Is(http.StatusUnauthorized) {
		return nil, ErrInvalidToken
	}

	var data ListPairedDevicesResponse
	if err := resp.Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
