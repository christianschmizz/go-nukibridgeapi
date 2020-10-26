package bridge_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func TestConnection_LockAction(t *testing.T) {
	conn, err := bridge.ConnectWithToken("192.168.1.11:8080", "abcdef")
	assert.NoError(t, err)

	t.Run("", func(t *testing.T) {
		_, err := conn.LockAction(nuki.NukiID{1, nuki.SmartLock}, nuki.OpenerLockActionActivateContinuousMode, bridge.NoWait())
		assert.Error(t, err)
	})
}

func TestConnection_LockState(t *testing.T) {
	conn, err := bridge.ConnectWithToken("192.168.1.11:8080", "abcdef")
	assert.NoError(t, err)

	t.Run("", func(t *testing.T) {
		_, err := conn.LockState(nuki.NukiID{1, nuki.SmartLock})
		assert.Error(t, err)
	})
}
