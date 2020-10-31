package bridgeapi_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

func TestDecode_Info(t *testing.T) {
	infoJSON, err := ioutil.ReadFile(filepath.Join("testdata", "info.json"))
	assert.NoError(t, err)

	var info bridgeapi.InfoResponse
	err = json.Unmarshal(infoJSON, &info)
	assert.NoError(t, err)

	ts, err := time.Parse(time.RFC3339, "2020-10-26T22:50:56+00:00")
	assert.NoError(t, err)

	assert.Equal(t, bridgeapi.InfoResponse{
		BridgeType: bridgeapi.TypeHardware,
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
		ScanResults: []bridgeapi.ScanResult{
			{509600314, 2, "Nuki_Opener_1E5FE23A", -57, true},
			{597878773, 0, "Nuki_23A2E7F5", -61, true},
		},
	}, info)
}