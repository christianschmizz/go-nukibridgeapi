// +build integration

package bridgeapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

func TestConnection_Info(t *testing.T) {
	conn := bridgeConn(t, *host, *token)
	info, err := conn.Info()
	if assert.NoError(t, err) {
		assert.Equal(t, bridgeapi.TypeHardware, info.BridgeType)
		assert.Len(t, info.ScanResults, 2)
	}
}
