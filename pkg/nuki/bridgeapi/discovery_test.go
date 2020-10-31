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

func TestDecode_Discover(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("testdata", "discover.json"))
	assert.NoError(t, err)

	{
		var discovery bridgeapi.DiscoverResponse
		err := json.Unmarshal(data, &discovery)
		assert.NoError(t, err)
		updated, _ := time.Parse(time.RFC3339, "2020-10-24T17:47:13Z")
		assert.Equal(t, bridgeapi.DiscoverResponse{
			[]bridgeapi.BridgeInfo{
				{
					BridgeID:    448942400,
					IP:          "192.168.1.50",
					Port:        8080,
					DateUpdated: updated,
				},
			},
			0,
		}, discovery)
	}
}
