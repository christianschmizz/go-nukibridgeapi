package bridgeapi_test

import (
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/internal/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

func Test_ListPairedDevicesResponse(t *testing.T) {
	t.Run("sucessful list", func(t *testing.T) {
		conn, err := bridgeapi.ConnectWithToken("127.0.0.1:8080", "abcdef",
			bridgeapi.UseClient(mocks.NewFileResponseClient(filepath.Join("testdata", "list.json"), http.StatusOK)))
		assert.NoError(t, err)

		reso, err := conn.ListPairedDevices()
		assert.NoError(t, err)

		ts, err := time.Parse(time.RFC3339, "2020-01-30T20:00:00+00:00")
		assert.NoError(t, err)

		expectedResponse := bridgeapi.ListPairedDevicesResponse{
			{ID: 527875674, Type: 0, Name: "Eine Tür", FirmwareVersion: "2.8.15",
				LastKnownState: bridgeapi.LastKnownState{
					Mode: 2, State: 3, StateName: "unlocked", BatteryCritical: false,
					BatteryCharging: false, BatteryChargeState: 100, KeypadBatteryCritical: false,
					DoorsensorState: 2, DoorsensorStateName: "door closed", RingactionState: false,
					RingactionTimestamp: time.Time{}, Timestamp: ts},
			},
			{ID: 519611324, Type: 2, Name: "Andere Tür", FirmwareVersion: "1.5.3",
				LastKnownState: bridgeapi.LastKnownState{
					Mode: 2, State: 1, StateName: "online", BatteryCritical: false,
					BatteryCharging: false, BatteryChargeState: 0, KeypadBatteryCritical: false,
					DoorsensorState: 0, DoorsensorStateName: "", RingactionState: false,
					RingactionTimestamp: ts, Timestamp: ts},
			},
		}
		assert.Equal(t, expectedResponse, reso)
	})

	t.Run("failed auth", func(t *testing.T) {
		conn, err := bridgeapi.ConnectWithToken("127.0.0.1:8080", "abcdef",
			bridgeapi.UseClient(mocks.NewFileResponseClient(filepath.Join("testdata", "list.json"), http.StatusUnauthorized)))
		assert.NoError(t, err)

		_, err = conn.ListPairedDevices()
		assert.Error(t, err)
		assert.IsType(t, bridgeapi.ErrInvalidToken, err)

	})
}
