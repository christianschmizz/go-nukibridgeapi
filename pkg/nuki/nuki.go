package nuki

import (
	"fmt"
	"strings"

	"github.com/stretchr/stew/slice"
)

// LockState describes the state of a device
type LockState int

// LockAction describes an action to be executed on a device
type LockAction int

// Int retrieves int value of the action
func (t LockAction) Int() int {
	return int(t)
}

// DeviceType describes the type of a device
type DeviceType int

// Int retrieves int value of the type
func (t DeviceType) Int() int {
	return int(t)
}

// LockMode describes the operational mode of a device
type LockMode int

// DoorsensorState describes the state of a doorsensor
type DoorsensorState int

// The ID is a unique identifier over all Nuki devices. The numeric ID is
// not unique between products, e.g. a smart lock's ID may collide with an
// opener ID.
type ID struct {
	DeviceID   int
	DeviceType DeviceType
}

const (
	// SmartLock is a Nuki Smart Lock
	SmartLock DeviceType = 0

	// Opener is a Nuki Opener used for integration into door-opener systems
	Opener DeviceType = 2

	// SmartDoor is a Nuki Smart Door
	SmartDoor DeviceType = 3

	// SmartLock3 is a Nuki Smart Lock 3.0 (Pro)
	SmartLock3 DeviceType = 4
)

const (
	// DoorMode is the operational mode after complete setup (applies to smart lock and opener)
	DoorMode LockMode = 2

	// ContinuousMode means Ring to Open is permanently active (applies to opener only)
	ContinuousMode LockMode = 3
)

const (
	// SmartLockStateUncalibrated means the smart lock is not yet configured
	SmartLockStateUncalibrated LockState = 0

	// SmartLockStateLocked means the door at the smart lock is/should be locked
	SmartLockStateLocked LockState = 1

	// SmartLockStateUnlocking means the smart lock is currently in the progress of unlocking the door
	SmartLockStateUnlocking LockState = 2

	// SmartLockStateUnlocked means the door at the smart lock is currently unlocked
	SmartLockStateUnlocked LockState = 3

	// SmartLockStateLocking means the smart lock is currently in the progress of locking the door
	SmartLockStateLocking LockState = 4

	// SmartLockStateUnlatched means the door at the smart lock us currently not closed
	SmartLockStateUnlatched LockState = 5

	// SmartLockStateUnlockedLockAndGo means the door is currently unlocked and Lock and Go is active
	SmartLockStateUnlockedLockAndGo LockState = 6

	// SmartLockStateUnlatching means the smart lock is currently in the progress of unlatching the door
	SmartLockStateUnlatching LockState = 7

	// SmartLockStateMotorBlocked means the smart lock cannot rotate the key anymore
	SmartLockStateMotorBlocked LockState = 254

	// SmartLockStateUndefined means the smart lock is in an undefined/unknown state
	SmartLockStateUndefined LockState = 255
)

const (
	// OpenerStateUntrained means the opener is not yet configured
	OpenerStateUntrained LockState = 0

	// OpenerStateOnline means that the opener is connected to the bridge and idling
	OpenerStateOnline LockState = 1

	// OpenerStateRtoActive means that Ring To Open is active at the opener
	OpenerStateRtoActive LockState = 3

	// OpenerStateOpen - meaning unknown
	OpenerStateOpen LockState = 5

	// OpenerStateOpening means that the opener is currently opening
	OpenerStateOpening LockState = 7

	// OpenerStateBootRun - meaning unknown
	OpenerStateBootRun LockState = 253

	// OpenerStateUndefined means that the state of the opener is unknown.
	OpenerStateUndefined LockState = 255
)

const (
	// SmartLockActionUnlock describes the action to unlock a smart lock
	SmartLockActionUnlock LockAction = 1

	// SmartLockActionLock describes the action to lock a smart lock
	SmartLockActionLock LockAction = 2

	// SmartLockActionUnlatch describes the action to open the door
	SmartLockActionUnlatch LockAction = 3

	// SmartLockActionLockAndGo describes the action to enable Lock and Go at a smart lock
	SmartLockActionLockAndGo LockAction = 4

	// SmartLockActionLockAndGoWithUnlatch describes the action to enable Lock and Go with unlatch at a smart lock
	SmartLockActionLockAndGoWithUnlatch LockAction = 5
)

const (

	// OpenerLockActionActivateRto describes the action to enable Ring To Open at the opener
	OpenerLockActionActivateRto LockAction = 1

	// OpenerLockActionDeactivateRto describes the action to disable Ring To Open at the opener
	OpenerLockActionDeactivateRto LockAction = 2

	// OpenerLockActionElectricStrikeActuation describes the action to enable electric strike actuation at the opener
	OpenerLockActionElectricStrikeActuation LockAction = 3

	// OpenerLockActionActivateContinuousMode describes the action to enable activate continuous mode at the opener
	OpenerLockActionActivateContinuousMode LockAction = 4

	// OpenerLockActionDeactivateContinuousMode describes the action to deactivate continuous mode at the opener
	OpenerLockActionDeactivateContinuousMode LockAction = 5
)

const (
	// DoorsensorStateDeactivated describes that the doorsensor is deativated
	DoorsensorStateDeactivated DoorsensorState = 1

	// DoorsensorStateDoorClosed describes that the door was detected to be closed
	DoorsensorStateDoorClosed DoorsensorState = 2

	// DoorsensorStateDoorOpened describes that the door was deteced to be open
	DoorsensorStateDoorOpened DoorsensorState = 3

	// DoorsensorStateDoorStateUnknown describes that the state of the door is unknown
	DoorsensorStateDoorStateUnknown DoorsensorState = 4

	// DoorsensorStateCalibrating describes that the doorsensor is currently being calibrated
	DoorsensorStateCalibrating DoorsensorState = 5

	// DoorsensorStateUncalibrated indicates an uncalibrated sensor
	DoorsensorStateUncalibrated DoorsensorState = 16

	DoorsensorStateRemoved DoorsensorState = 240

	DoorsensorStateUnknown DoorsensorState = 255
)

// IsActionSupportedByDeviceType checks whether the given action is supported by that device type
func IsActionSupportedByDeviceType(action LockAction, deviceType DeviceType) bool {
	if (deviceType == SmartLock || deviceType == SmartLock3) && slice.Contains(SmartLockActions, action) {
		return true
	} else if deviceType == Opener && slice.Contains(OpenerLockActions, action) {
		return true
	} else {
		return false
	}
}

// SmartLockActionFromString retrieves the appropriate action for a smart lock from a string
func SmartLockActionFromString(s string) (e LockAction, err error) {
	switch strings.ToLower(s) {
	case "unlock":
		e = SmartLockActionUnlock
	case "lock":
		e = SmartLockActionLock
	case "unlatch":
		e = SmartLockActionUnlatch
	case "lockandgo":
		e = SmartLockActionLockAndGo
	case "lockandgowithunlatch":
		e = SmartLockActionLockAndGoWithUnlatch
	default:
		err = fmt.Errorf("unknown smart lock action: %s", s)
	}
	return
}

// OpenerLockActionFromString retrieves the appropriate action for an opener from a string
func OpenerLockActionFromString(s string) (e LockAction, err error) {
	switch strings.ToLower(s) {
	case "rto_on":
		e = OpenerLockActionActivateRto
	case "rto_off":
		e = OpenerLockActionDeactivateRto
	case "esa":
		e = OpenerLockActionElectricStrikeActuation
	case "cm_on":
		e = OpenerLockActionActivateContinuousMode
	case "cm_off":
		e = OpenerLockActionDeactivateContinuousMode
	default:
		err = fmt.Errorf("unknown opener action: %s", s)
	}
	return
}

// LockActionFromString retrieves the appropriate action for a device type from a string
func LockActionFromString(s string, t DeviceType) (action LockAction, err error) {
	switch t {
	case SmartLock:
		action, err = SmartLockActionFromString(s)
	case Opener:
		action, err = OpenerLockActionFromString(s)
	default:
		err = fmt.Errorf("unknown device type: %s", s)
	}
	return
}

var (
	// SmartLockActions contains all actions available for a smart lock
	SmartLockActions = [...]LockAction{
		SmartLockActionUnlock,
		SmartLockActionLock,
		SmartLockActionUnlatch,
		SmartLockActionLockAndGo,
		SmartLockActionLockAndGoWithUnlatch,
	}

	// OpenerLockActions contains all actions available for an opener
	OpenerLockActions = [...]LockAction{
		OpenerLockActionActivateRto,
		OpenerLockActionDeactivateRto,
		OpenerLockActionElectricStrikeActuation,
		OpenerLockActionActivateContinuousMode,
		OpenerLockActionDeactivateContinuousMode,
	}

	// SmartLockStates contains all valid states of a smart lock
	SmartLockStates = [...]LockState{
		SmartLockStateUncalibrated,
		SmartLockStateLocked,
		SmartLockStateUnlocking,
		SmartLockStateUnlocked,
		SmartLockStateLocking,
		SmartLockStateUnlatched,
		SmartLockStateUnlockedLockAndGo,
		SmartLockStateUnlatching,
		SmartLockStateMotorBlocked,
		SmartLockStateUndefined,
	}

	// OpenerLockStates contains all valid states of an opener
	OpenerLockStates = [...]LockState{
		OpenerStateUntrained,
		OpenerStateOnline,
		OpenerStateRtoActive,
		OpenerStateOpen,
		OpenerStateOpening,
		OpenerStateBootRun,
		OpenerStateUndefined,
	}
)
