package bridgeapi

import (
	"fmt"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

// LockResponse represents the result of an locking request
type LockResponse struct {
	Success         bool `json:"success"`
	BatteryCritical bool `json:"batteryCritical"`
}

type lockOptions struct {
	DeviceID   int             `url:"nukiId"`
	DeviceType nuki.DeviceType `url:"deviceType"`
}

// Lock sends a simple lock action "lock" to the given device
func (c *Connection) Lock(nukiID nuki.NukiID, options ...func(*lockOptions)) (*LockResponse, error) {
	o := &lockOptions{nukiID.DeviceID, nukiID.DeviceType}
	for _, opt := range options {
		opt(o)
	}

	var response LockResponse
	if err := c.get(c.hashedURL("lock", o), &response); err != nil {
		return nil, fmt.Errorf("could not execute lock: %w", err)
	}
	return &response, nil
}
