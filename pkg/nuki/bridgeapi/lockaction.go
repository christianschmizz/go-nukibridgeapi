package bridgeapi

import (
	"fmt"

	"github.com/stretchr/stew/slice"

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

func isActionSupported(action nuki.LockAction, deviceType nuki.DeviceType) bool {
	if deviceType == nuki.SmartLock && slice.Contains(nuki.SmartLockActions, action) {
		return true
	} else if deviceType == nuki.Opener && slice.Contains(nuki.OpenerLockActions, action) {
		return true
	} else {
		return false
	}
}

// LockAction performs a action on the device with the given ID.
func (c *Connection) LockAction(nukiID nuki.ID, action nuki.LockAction, options ...func(*lockActionOptions)) (*LockActionResponse, error) {
	if !isActionSupported(action, nukiID.DeviceType) {
		return nil, fmt.Errorf("unsupported lockAction: %v", action)
	}

	o := &lockActionOptions{nukiID.DeviceID, nukiID.DeviceType, action.Int(), true}
	for _, opt := range options {
		opt(o)
	}

	var response LockActionResponse
	if err := c.get(c.hashedURL("lockAction", o), &response); err != nil {
		return nil, fmt.Errorf("could not execute lockAction: %w", err)
	}
	return &response, nil
}
