// +build integration

package bridgeapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSession_ListPairedDevices(t *testing.T) {
	conn := bridgeConn(t)
	devices, err := conn.ListPairedDevices()
	if assert.NoError(t, err) {
		assert.Len(t, devices, 2)
	}
}
