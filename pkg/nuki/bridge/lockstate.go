package bridge

import (
	"fmt"
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

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

// Retrieves and returns the current lock state of a given Nuki device
func (c *Connection) LockState(nukiID nuki.NukiID) (*LockStateResponse, error) {
	options := &lockStateOptions{nukiID.DeviceID, nukiID.DeviceType}

	var response LockStateResponse
	if err := c.get(c.hashedURL("lockState", options), &response); err != nil {
		return nil, fmt.Errorf("could not fetch lockState: %w", err)
	}
	return &response, nil
}
