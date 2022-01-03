package bridgeapi

import (
	"net/http"
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

// LockStateResponse contains the response results from the respective API call
type LockStateResponse struct {
	Mode                  nuki.LockMode        `json:"mode"`
	State                 nuki.LockState       `json:"state"`
	StateName             string               `json:"stateName"`
	BatteryCritical       bool                 `json:"batteryCritical"`
	BatteryCharging       bool                 `json:"batteryCharging"`
	BatteryChargeState    uint8                `json:"batteryChargeState"`
	KeypadBatteryCritical bool                 `json:"keypadBatteryCritical"`
	DoorsensorState       nuki.DoorsensorState `json:"doorsensorState,omitempty"`
	DoorsensorStateName   string               `json:"doorsensorStateName,omitempty"`
	RingactionState       bool                 `json:"ringactionState,omitempty"`
	RingactionTimestamp   time.Time            `json:"ringactionTimestamp,omitempty"`
	Success               bool                 `json:"success"`
}

type lockStateOptions struct {
	DeviceID   int             `url:"nukiId"`
	DeviceType nuki.DeviceType `url:"deviceType"`
}

// LockState retrieves the current state of the given device
func (c *Connection) LockState(nukiID nuki.ID) (*LockStateResponse, error) {
	options := &lockStateOptions{
		DeviceID:   nukiID.DeviceID,
		DeviceType: nukiID.DeviceType,
	}

	resp, err := c.request("lockState", options)
	if err != nil {
		return nil, err
	}

	if resp.Is(http.StatusUnauthorized) {
		return nil, ErrInvalidToken
	} else if resp.Is(http.StatusNotFound) {
		return nil, ErrUnknownDevice
	} else if resp.Is(http.StatusServiceUnavailable) {
		return nil, ErrDeviceIsTemporarilyOffline
	}

	var data LockStateResponse
	if err := resp.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
