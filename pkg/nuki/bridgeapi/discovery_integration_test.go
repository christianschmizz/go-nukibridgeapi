// +build integration

package bridgeapi_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

func TestDiscover(t *testing.T) {
	d, err := bridgeapi.Discover()
	if assert.NoError(t, err) {
		assert.Equal(t, 0, d.ErrorCode)
	}
}
