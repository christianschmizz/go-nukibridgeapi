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

func Test_Log(t *testing.T) {
	t.Run("successful log", func(t *testing.T) {
		conn, err := bridgeapi.ConnectWithToken("127.0.0.1:8080", "abcdef",
			bridgeapi.UseClient(mocks.NewFileResponseClient(filepath.Join("testdata", "log.json"), http.StatusOK)))
		assert.NoError(t, err)

		resp, err := conn.Log(0, 1000)
		assert.NoError(t, err)
		assert.Len(t, resp, 100)

		t.Run("SSE-PushNukisRequest", func(t *testing.T) {
			ts, err := time.Parse(time.RFC3339, "2020-11-01T21:53:37+00:00")
			assert.NoError(t, err)

			expectedResponse := bridgeapi.LogEntry{
				Timestamp: ts,
				Type:      "SSE-PushNukisRequest",
				Count:     2,
			}
			assert.Equal(t, expectedResponse, resp[2])
		})

		t.Run("BLE-Disconnected", func(t *testing.T) {
			ts, err := time.Parse(time.RFC3339, "2020-11-01T21:53:34+00:00")
			assert.NoError(t, err)

			expectedResponse := bridgeapi.LogEntry{
				Timestamp: ts,
				ID:        "1E5FE23A",
				Type:      "BLE-Disconnected",
				PairIndex: 1,
				BleHandle: "0001",
			}
			assert.Equal(t, expectedResponse, resp[3])
		})

		t.Run("BLE-Connect", func(t *testing.T) {
			ts, err := time.Parse(time.RFC3339, "2020-11-01T21:53:33+00:00")
			assert.NoError(t, err)

			expectedResponse := bridgeapi.LogEntry{
				Timestamp: ts,
				Type:      "BLE-Connect",
				MacAddr:   "54D2725FE23A",
			}
			assert.Equal(t, expectedResponse, resp[12])
		})
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
