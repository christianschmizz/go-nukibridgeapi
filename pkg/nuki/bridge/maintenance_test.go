package bridge_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func TestConnection_Log(t *testing.T) {
	conn, err := bridge.ConnectWithToken("192.168.1.11:8080", "abcdef")
	assert.NoError(t, err)

	t.Run("", func(t *testing.T) {
		logs, err := conn.Log(0, 1000)
		assert.NoError(t, err)
		for _, l := range logs {
			fmt.Printf("%+v\n", l)
		}
	})
}
