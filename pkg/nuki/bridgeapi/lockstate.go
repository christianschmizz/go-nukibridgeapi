package bridgeapi

import (
	"fmt"
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

// LockStateResponse contains the response results from the respective API call
type LockStateResponse struct {
	Mode                  int       `json:"mode"`
	State                 int       `json:"state"`
	StateName             string    `json:"stateName"`
	BatteryCritical       bool      `json:"batteryCritical"`
	KeypadBatteryCritical bool      `json:"keypadBatteryCritical"`
	DoorsensorState       int       `json:"doorsensorState,omitempty"`
	DoorsensorStateName   string    `json:"doorsensorStateName,omitempty"`
	RingactionState       bool      `json:"ringactionState,omitempty"`
	RingactionTimestamp   time.Time `json:"ringactionTimestamp,omitempty"`
	Success               bool      `json:"success"`
}

type lockStateOptions struct {
	DeviceID   int             `url:"nukiId"`
	DeviceType nuki.DeviceType `url:"deviceType"`
}

// LockState retrieves the current state of the given device
func (c *Connection) LockState(nukiID nuki.ID) (*LockStateResponse, error) {
	options := &lockStateOptions{nukiID.DeviceID, nukiID.DeviceType}

	var response LockStateResponse
	if err := c.get(c.hashedURL("lockState", options), &response); err != nil {
		return nil, fmt.Errorf("could not fetch lockState: %w", err)
	}
	return &response, nil
}
