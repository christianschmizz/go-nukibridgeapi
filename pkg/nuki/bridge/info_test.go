package bridge_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	nukibridge "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func TestConnection_Info(t *testing.T) {
	conn, err := nukibridge.ConnectWithToken("192.168.1.11:8080", "abcdef")
	if assert.NoError(t, err) {
		info, err := conn.Info()
		if assert.NoError(t, err) {
			assert.Equal(t, nukibridge.Info{
				BridgeType: nukibridge.TypeHardware,
				IDs: struct {
					HardwareID int `json:"hardwareId"`
					ServerID   int `json:"serverId"`
				}{548263954, 448942400},
				Versions: struct {
					FirmwareVersion     string `json:"firmwareVersion"`
					WifiFirmwareVersion string `json:"wifiFirmwareVersion"`
				}{"2.7.0", "2.1.17"},
				Uptime:          2067,
				CurrentTime:     time.Now().UTC(),
				ServerConnected: true,
				ScanResults: []nukibridge.ScanResult{
					{509600314, 0, "Nuki_Opener_1E5FE23A", -57, true},
				},
			}, *info)
		}
	}
}

func TestSession_DecodeInfo(t *testing.T) {
	infoJson, err := ioutil.ReadFile(filepath.Join("testdata", "info.json"))
	assert.NoError(t, err)

	var info nukibridge.Info
	err = json.Unmarshal(infoJson, &info)
	assert.NoError(t, err)

	ts, err := time.Parse(time.RFC3339, "2020-10-26T22:50:56+00:00")
	assert.NoError(t, err)

	assert.Equal(t, nukibridge.Info{
		BridgeType: nukibridge.TypeHardware,
		IDs: struct {
			HardwareID int `json:"hardwareId"`
			ServerID   int `json:"serverId"`
		}{548263954, 448942400},
		Versions: struct {
			FirmwareVersion     string `json:"firmwareVersion"`
			WifiFirmwareVersion string `json:"wifiFirmwareVersion"`
		}{"2.7.0", "2.1.17"},
		Uptime:          1278,
		CurrentTime:     ts,
		ServerConnected: true,
		ScanResults: []nukibridge.ScanResult{
			{509600314, 2, "Nuki_Opener_1E5FE23A", -57, true},
			{597878773, 0, "Nuki_23A2E7F5", -61, true},
		},
	}, info)
}
