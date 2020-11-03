package bridgeapi

import (
	"net/http"

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
	o := &unlockOptions{
		DeviceID:   nukiID.DeviceID,
		DeviceType: nukiID.DeviceType,
	}
	for _, opt := range options {
		opt(o)
	}

	resp, err := c.request("unlock", nil)
	if err != nil {
		return nil, err
	}

	if resp.Is(http.StatusUnauthorized) {
		return nil, ErrInvalidToken
	} else if resp.Is(http.StatusNotFound) {
		return nil, ErrUnknownDevice
	} else if resp.Is(http.StatusServiceUnavailable) {
		return nil, ErrDeviceOffline
	}

	var data UnlockResponse
	if err := resp.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
