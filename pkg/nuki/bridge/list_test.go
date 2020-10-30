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

func TestDecode_ListPairedDevicesResponse(t *testing.T) {
	listJson, err := ioutil.ReadFile(filepath.Join("testdata", "list.json"))
	assert.NoError(t, err)

	var reso nukibridge.ListPairedDevicesResponse
	err = json.Unmarshal(listJson, &reso)
	assert.NoError(t, err)

	ts, err := time.Parse(time.RFC3339, "2020-01-30T20:00:00+00:00")
	assert.NoError(t, err)

	expectedResponse := nukibridge.ListPairedDevicesResponse{
		{ID: 527875674, Type: 0, Name: "Eine Tür", LastKnownState: nukibridge.LastKnownState{
			Mode: 2, State: 3, StateName: "unlocked", BatteryCritical: false,
			BatteryCharging: false, BatteryChargeState: 100, KeypadBatteryCritical: false,
			DoorsensorState: 2, DoorsensorStateName: "door closed", RingactionState: false,
			RingactionTimestamp: time.Time{}, Timestamp: ts},
		},
		{ID: 519611324, Type: 2, Name: "Andere Tür", LastKnownState: nukibridge.LastKnownState{
			Mode: 2, State: 1, StateName: "online", BatteryCritical: false,
			BatteryCharging: false, BatteryChargeState: 0, KeypadBatteryCritical: false,
			DoorsensorState: 0, DoorsensorStateName: "", RingactionState: false,
			RingactionTimestamp: ts, Timestamp: ts},
		},
	}
	assert.Equal(t, expectedResponse, reso)
}