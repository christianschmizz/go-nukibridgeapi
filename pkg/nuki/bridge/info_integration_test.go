// +build integration

package bridge_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nukibridge "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func TestConnection_Info(t *testing.T) {
	conn := bridgeConn(t, *host, *token)
	info, err := conn.Info()
	if assert.NoError(t, err) {
		assert.Equal(t, nukibridge.TypeHardware, info.BridgeType)
		assert.Len(t, info.ScanResults, 2)
	}
}
