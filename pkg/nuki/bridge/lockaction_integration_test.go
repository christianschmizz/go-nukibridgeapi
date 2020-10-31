// +build integration

package bridge_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func TestConnection_LockAction(t *testing.T) {
	conn := bridgeConn(t, *host, *token)

	var (
		err error
		i   *bridge.InfoResponse
		r   bridge.ScanResult
	)

	t.Run("fetch info", func(t *testing.T) {
		i, err = conn.Info()
		assert.NoError(t, err)
		assert.Len(t, i.ScanResults, 2)
		r = i.ScanResults[0]
	})

	t.Run("deactivate rto", func(t *testing.T) {
		result, err := conn.LockAction(*r.NukiID(), nuki.OpenerLockActionDeactivateRto, bridge.Wait())
		assert.NoError(t, err)
		assert.True(t, result.Success)
	})
}

func TestConnection_LockState(t *testing.T) {
	conn := bridgeConn(t, *host, *token)

	var (
		err     error
		devices bridge.ListPairedDevicesResponse
	)

	t.Run("fetch list of paired devices", func(t *testing.T) {
		devices, err = conn.ListPairedDevices()
		assert.NoError(t, err)
		assert.Len(t, devices, 2)
	})

	t.Run("read lock's state", func(t *testing.T) {
		id := nuki.NukiID{DeviceID: devices[0].ID, DeviceType: nuki.DeviceType(devices[0].Type)}
		result, err := conn.LockState(id)
		assert.NoError(t, err)
		assert.True(t, result.Success)
	})
}
