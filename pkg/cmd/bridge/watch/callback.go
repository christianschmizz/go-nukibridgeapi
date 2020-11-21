package watch

import (
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

type CallbackData struct {
	ID                    int                  `json:"nukiId"`
	Type                  nuki.DeviceType      `json:"deviceType"`
	Mode                  nuki.LockMode        `json:"mode"`
	State                 nuki.LockState       `json:"state"`
	StateName             string               `json:"stateName"`
	BatteryCritical       bool                 `json:"batteryCritical"`
	KeypadBatteryCritical bool                 `json:"keypadBatteryCritical"`
	DoorsensorState       nuki.DoorsensorState `json:"doorsensorState,omitempty"`
	DoorsensorStateName   string               `json:"doorsensorStateName,omitempty"`
	RingactionTimestamp   time.Time            `json:"ringactionTimestamp,omitempty" dbus:"-"` // Encoding of timestamps leads to an error '"dbus: connection closed by user"'
	RingactionState       bool                 `json:"ringactionState,omitempty"`
}

// NukiID assembles the ID from a result
func (d *CallbackData) NukiID() *nuki.ID {
	return &nuki.ID{
		DeviceID:   d.ID,
		DeviceType: d.Type,
	}
}
