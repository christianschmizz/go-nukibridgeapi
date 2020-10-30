// +build integration

package bridge_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	nukibridge "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func TestDiscover(t *testing.T) {
	d, err := nukibridge.Discover()
	if assert.NoError(t, err) {
		assert.Equal(t, 0, d.ErrorCode)
	}
}
