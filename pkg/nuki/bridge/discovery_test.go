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

func TestDiscover(t *testing.T) {
	d, err := nukibridge.Discover()
	if assert.NoError(t, err) {
		assert.Equal(t, 0, d.ErrorCode)
	}
}

func TestDiscover_Decode(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("testdata", "discover.json"))
	assert.NoError(t, err)

	{
		var discovery nukibridge.Discovery
		err := json.Unmarshal(data, &discovery)
		assert.NoError(t, err)
		updated, _ := time.Parse(time.RFC3339, "2020-10-24T17:47:13Z")
		assert.Equal(t, nukibridge.Discovery{
			[]nukibridge.BridgeInfo{{
				BridgeID:    448942400,
				IP:          "192.168.1.50",
				Port:        8080,
				DateUpdated: updated,
			}},
			0,
		}, discovery)
	}
}
