package bridgeapi

import (
	"fmt"
	"net/http"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

// LockActionResponse represents the result of a request to /lockAction
type LockActionResponse struct {
	Success         bool `json:"success"`
	BatteryCritical bool `json:"batteryCritical"`
}

type lockActionOptions struct {
	DeviceID   int             `url:"nukiId"`
	DeviceType nuki.DeviceType `url:"deviceType"`
	Action     int             `url:"action"`
	NoWait     bool            `url:"nowait"`
}

// Wait option is used for synchronously calling the API
func Wait() func(*lockActionOptions) {
	return func(options *lockActionOptions) {
		options.NoWait = false
	}
}

// LockAction performs a action on the device with the given ID.
func (c *Connection) LockAction(nukiID nuki.ID, action nuki.LockAction, options ...func(*lockActionOptions)) (*LockActionResponse, error) {
	if !nuki.IsActionSupportedByDeviceType(action, nukiID.DeviceType) {
		return nil, fmt.Errorf("unsupported lockAction: %v", action)
	}

	o := &lockActionOptions{
		DeviceID:   nukiID.DeviceID,
		DeviceType: nukiID.DeviceType,
		Action:     action.Int(),
		NoWait:     true,
	}
	for _, opt := range options {
		opt(o)
	}

	resp, err := c.request("lockAction", o)
	if err != nil {
		return nil, err
	}

	if resp.Is(http.StatusBadRequest) {
		return nil, &ErrInvalidAction{Action: action}
	} else if resp.Is(http.StatusUnauthorized) {
		return nil, ErrInvalidToken
	} else if resp.Is(http.StatusNotFound) {
		return nil, ErrUnknownDevice
	} else if resp.Is(http.StatusServiceUnavailable) {
		return nil, ErrDeviceIsTemporarilyOffline
	}

	var data LockActionResponse
	if err := resp.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
