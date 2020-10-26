package bridge

import (
	"fmt"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

type unlockResponse struct {
	Success         bool `json:"success"`
	BatteryCritical bool `json:"batteryCritical"`
}

type unlockOptions struct {
	DeviceID   int             `url:"nukiId"`
	DeviceType nuki.DeviceType `url:"deviceType"`
}

// Send the simple lock action "lock" to a given Nuki device
func (c *connection) Unlock(nukiID nuki.NukiID, options ...func(*unlockOptions)) (*unlockResponse, error) {
	o := &unlockOptions{nukiID.DeviceID, nukiID.DeviceType}
	for _, opt := range options {
		opt(o)
	}

	var response unlockResponse
	if err := c.get(c.hashedURL("unlock", o), &response); err != nil {
		return nil, fmt.Errorf("could not execute lock: %w", err)
	}
	return &response, nil
}
