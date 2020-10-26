package bridge_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	nukibridge "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func ExampleConnectWithToken() {
	conn, err := nukibridge.ConnectWithToken("192.168.1.11:8080", "abcdef")
	if err != nil {
		panic(err)
	}
	info, err := conn.Info()
	if err != nil {
		panic(err)
	}
	for _, result := range info.ScanResults {
		fmt.Println(result.Name)
	}
}

func ExampleScanOnConnect() {
	_, err := nukibridge.ConnectWithToken("192.168.1.11:8080", "abcdef", nukibridge.ScanOnConnect())
	if err != nil {
		panic(err)
	}
}

func TestSession_ListPairedDevices(t *testing.T) {
	conn, err := nukibridge.ConnectWithToken("192.168.1.11:8080", "abcdef")
	if assert.NoError(t, err) {
		devices, err := conn.ListPairedDevices()
		if assert.NoError(t, err) {
			assert.Len(t, devices, 2)
		}
	}
}
