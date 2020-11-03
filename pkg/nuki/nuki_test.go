package nuki_test

import (
	"testing"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
	"github.com/stretchr/testify/assert"
)

func TestLockActionFromString(t *testing.T) {
	tests := []struct {
		Name         string
		DeviceType   nuki.DeviceType
		Action       nuki.LockAction
		ExpectsError bool
	}{
		{"unlock", nuki.SmartLock, nuki.SmartLockActionUnlock, false},
		{"unLOCK", nuki.SmartLock, nuki.SmartLockActionUnlock, false},
		{"lock", nuki.SmartLock, nuki.SmartLockActionLock, false},
		{"LOCK", nuki.SmartLock, nuki.SmartLockActionLock, false},
		{"unlatch", nuki.SmartLock, nuki.SmartLockActionUnlatch, false},
		{"lockandgo", nuki.SmartLock, nuki.SmartLockActionLockAndGo, false},
		{"lockandgowithunlatch", nuki.SmartLock, nuki.SmartLockActionLockAndGoWithUnlatch, false},
		{"lockAndGoWithUnlatch", nuki.SmartLock, nuki.SmartLockActionLockAndGoWithUnlatch, false},
		{"", nuki.SmartLock, nuki.SmartLockActionLockAndGoWithUnlatch, true},
		{"Foo", nuki.SmartLock, nuki.SmartLockActionLockAndGoWithUnlatch, true},
	}

	for _, tt := range tests {
		got, err := nuki.LockActionFromString(tt.Name, tt.DeviceType)
		if tt.ExpectsError {
			assert.Error(t, err)
		} else {
			assert.Equal(t, got, tt.Action)
		}
	}
}
