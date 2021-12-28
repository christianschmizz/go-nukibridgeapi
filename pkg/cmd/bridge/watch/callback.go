package watch

import (
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

// CallbackData represents all data delivered by the Nuki bridge in the callback
type CallbackData struct {
	ID                    int                  `json:"nukiId"`
	Type                  nuki.DeviceType      `json:"deviceType"`
	Mode                  nuki.LockMode        `json:"mode"`
	State                 nuki.LockState       `json:"state"`
	StateName             string               `json:"stateName"`
	BatteryCritical       bool                 `json:"batteryCritical"`
	BatteryCharging       bool                 `json:"batteryCharging"`
	BatteryChargeState    uint8                `json:"batteryChargeState"`
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
