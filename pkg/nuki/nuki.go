package nuki

import (
	"fmt"
	"strings"
)

type LockState int

type LockAction int

func (t LockAction) Int() int {
	return int(t)
}

type DeviceType int

func (t DeviceType) Int() int {
	return int(t)
}

type LockMode int

type DoorsensorState int

// The NukiID is a unique identifier over all Nuki devices. The numeric ID is
// not unique between products, e.g. a smart lock's ID may collide with an
// opener ID.
type NukiID struct {
	DeviceID   int
	DeviceType DeviceType
}

const (
	SmartLock DeviceType = 0
	Opener    DeviceType = 2

	// Operation mode after complete setup (applies to Smart lock and Opener
	DoorMode LockMode = 2

	// Ring to Open permanently active (applies to Opener only)
	ContinuousMode LockMode = 3

	// Possible lock states for smart locks
	SmartLockStateUncalibrated      LockState = 0
	SmartLockStateLocked            LockState = 1
	SmartLockStateUnlocking         LockState = 2
	SmartLockStateUnlocked          LockState = 3
	SmartLockStateLocking           LockState = 4
	SmartLockStateUnlatched         LockState = 5
	SmartLockStateUnlockedLockAndGo LockState = 6
	SmartLockStateUnlatching        LockState = 7
	SmartLockStateMotorBlocked      LockState = 254
	SmartLockStateUndefined         LockState = 255

	// Possible lock states for opener
	OpenerStateUntrained LockState = 0
	OpenerStateOnline    LockState = 1
	OpenerStateRtoActive LockState = 3
	OpenerStateOpen      LockState = 5
	OpenerStateOpening   LockState = 7
	OpenerStateBootRun   LockState = 253
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
