package bridgeapi

import (
	"net/http"

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
func (c *Connection) Lock(nukiID nuki.ID, options ...func(*lockOptions)) (*LockResponse, error) {
	o := &lockOptions{
		DeviceID:   nukiID.DeviceID,
		DeviceType: nukiID.DeviceType,
	}
	for _, opt := range options {
		opt(o)
	}

	resp, err := c.request("lock", nil)
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

	var data LockResponse
	if err := resp.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
