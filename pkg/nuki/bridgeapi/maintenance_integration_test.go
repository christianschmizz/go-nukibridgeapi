// +build integration

package bridgeapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection_Log(t *testing.T) {
	conn := bridgeConn(t, *host, *token)

	{
		logs, err := conn.Log(0, 12)
		assert.NoError(t, err)
		assert.Len(t, logs, 12)
	}
}
