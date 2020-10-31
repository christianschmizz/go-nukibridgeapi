package nuki

import (
	"fmt"
	"strings"
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

// The NukiID is a unique identifier over all Nuki devices. The numeric ID is
// not unique between products, e.g. a smart lock's ID may collide with an
// opener ID.
type NukiID struct {
	DeviceID   int
	DeviceType DeviceType
}

const (
	// SmartLock is a Nuki Smart Lock
	SmartLock DeviceType = 0

	// Opener is a Nuki Opener used for integration into door-opener systems
	Opener DeviceType = 2

	// DoorMode is the operational mode after complete setup (applies to smart lock and opener)
	DoorMode LockMode = 2

	// ContinuousMode means Ring to Open is permanently active (applies to opener only)
	ContinuousMode LockMode = 3

	// SmartLockStateUncalibrated means the smart lock is not yet configured
	SmartLockStateUncalibrated LockState = 0

	// SmartLockStateLocked means the door at the smart lock is/should be locked
	SmartLockStateLocked LockState = 1

	// SmartLockStateUnlocking means the smart lock is currently in the progress of unlocking the door
	SmartLockStateUnlocking LockState = 2

	// SmartLockStateUnlocked means the door at the smart lock is currently unlocked
	SmartLockStateUnlocked LockState = 3

	// SmartLockStateUnlocking means the smart lock is currently in the progress of locking the door
	SmartLockStateLocking LockState = 4

	// SmartLockStateUnlatched means the door at the smart lock us currently not closed
	SmartLockStateUnlatched LockState = 5

	// SmartLockStateUnlockedLockAndGo means the door is currently unlocked and Lock and Go is active
	SmartLockStateUnlockedLockAndGo LockState = 6

	// SmartLockStateUnlatched means the smart lock is currently in the progress of unlatching the door
	SmartLockStateUnlatching LockState = 7

	// SmartLockStateMotorBlocked means the smart lock cannot rotate the key anymore
	SmartLockStateMotorBlocked LockState = 254

	// SmartLockStateUndefined means the smart lock is in an undefined/unknown state
	SmartLockStateUndefined LockState = 255

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

	SmartLockActionUnlock               LockAction = 1
	SmartLockActionLock                 LockAction = 2
	SmartLockActionUnlatch              LockAction = 3
	SmartLockActionLockAndGo            LockAction = 4
	SmartLockActionLockAndGoWithUnlatch LockAction = 5

	OpenerLockActionActivateRto              LockAction = 1
	OpenerLockActionDeactivateRto            LockAction = 2
	OpenerLockActionElectricStrikeActuation  LockAction = 3
	OpenerLockActionActivateContinuousMode   LockAction = 4
	OpenerLockActionDeactivateContinuousMode LockAction = 5

	DoorsensorStateDeactivated      DoorsensorState = 1
	DoorsensorStateDoorClosed       DoorsensorState = 2
	DoorsensorStateDoorOpened       DoorsensorState = 3
	DoorsensorStateDoorStateUnknown DoorsensorState = 4
	DoorsensorStateCalibrating      DoorsensorState = 5
)

// SmartLockActionFromString retrieves the appropriate action from a string
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

var (
	SmartLockActions = [...]LockAction{
		SmartLockActionUnlock,
		SmartLockActionLock,
		SmartLockActionUnlatch,
		SmartLockActionLockAndGo,
		SmartLockActionLockAndGoWithUnlatch,
	}

	OpenerLockActions = [...]LockAction{
		OpenerLockActionActivateRto,
		OpenerLockActionDeactivateRto,
		OpenerLockActionElectricStrikeActuation,
		OpenerLockActionActivateContinuousMode,
		OpenerLockActionDeactivateContinuousMode,
	}

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
