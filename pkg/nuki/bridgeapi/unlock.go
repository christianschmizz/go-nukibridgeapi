package bridgeapi

import (
	"fmt"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

// UnlockResponse represents the result of a unlocking request
type UnlockResponse struct {
	Success         bool `json:"success"`
	BatteryCritical bool `json:"batteryCritical"`
}

type unlockOptions struct {
	DeviceID   int             `url:"nukiId"`
	DeviceType nuki.DeviceType `url:"deviceType"`
}

// Unlock sends a simple lock action "lock" to the given device
func (c *Connection) Unlock(nukiID nuki.ID, options ...func(*unlockOptions)) (*UnlockResponse, error) {
	o := &unlockOptions{nukiID.DeviceID, nukiID.DeviceType}
	for _, opt := range options {
		opt(o)
	}

	var response UnlockResponse
	if err := c.get(c.hashedURL("unlock", o), &response); err != nil {
		return nil, fmt.Errorf("could not execute lock: %w", err)
	}
	return &response, nil
}
