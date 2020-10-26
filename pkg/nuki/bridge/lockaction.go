package bridge

import (
	"fmt"

	"github.com/stretchr/stew/slice"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

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

func NoWait() func(*lockActionOptions) {
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

// Performs a lock action on the Nuki device with the given nukiID.
func (c *connection) LockAction(nukiID nuki.NukiID, action nuki.LockAction, options ...func(*lockActionOptions)) (*LockActionResponse, error) {
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
