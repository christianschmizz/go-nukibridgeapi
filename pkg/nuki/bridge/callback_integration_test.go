// +build integration

package bridge_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func Test_CallbackHandling(t *testing.T) {
	conn := bridgeConn(t, *host, *token)

	var (
		err                   error
		addCallbackResponse   *bridge.AddCallbackResponse
		listCallbacksResponse *bridge.ListCallbacksResponse
		firstCallback         bridge.Callback
		secondCallback        bridge.Callback
	)

	t.Run("Add first callback", func(t *testing.T) {
		addCallbackResponse, err = conn.AddCallback("http://192.168.1.1:8080")
		assert.NoError(t, err)
		assert.True(t, addCallbackResponse.Success)
	})

	t.Run("Check for first callback", func(t *testing.T) {
		listCallbacksResponse, err = conn.ListCallbacks()
		assert.NoError(t, err)
		assert.Len(t, listCallbacksResponse.Callbacks, 1)

		firstCallback = listCallbacksResponse.Callbacks[0]
		assert.Equal(t, "http://192.168.1.1:8080", firstCallback.URL)
	})

	t.Run("Add second callback", func(t *testing.T) {
		addCallbackResponse, err = conn.AddCallback("http://192.168.1.2:8080")
		assert.NoError(t, err)
		assert.True(t, addCallbackResponse.Success)
	})

	t.Run("Check for second callback", func(t *testing.T) {
		listCallbacksResponse, err = conn.ListCallbacks()
		assert.NoError(t, err)
		assert.Len(t, listCallbacksResponse.Callbacks, 2)

		secondCallback = listCallbacksResponse.Callbacks[1]
		assert.Equal(t, "http://192.168.1.2:8080", secondCallback.URL)
	})

	t.Run("Remove first callback", func(t *testing.T) {
		remove, err := conn.RemoveCallback(firstCallback.ID)
		assert.NoError(t, err)
		assert.True(t, remove.Success)
	})

	t.Run("Ensure removal and check for remaining callback", func(t *testing.T) {
		listCallbacksResponse, err = conn.ListCallbacks()
		assert.NoError(t, err)
		assert.Len(t, listCallbacksResponse.Callbacks, 1)

		callback := listCallbacksResponse.Callbacks[0]
		assert.Equal(t, "http://192.168.1.2:8080", callback.URL)
	})

	t.Run("Remove second callback", func(t *testing.T) {
		remove, err := conn.RemoveCallback(secondCallback.ID)
		assert.NoError(t, err)
		assert.True(t, remove.Success)
	})

	t.Run("Ensure no callbacks remain", func(t *testing.T) {
		listCallbacksResponse, err = conn.ListCallbacks()
		assert.NoError(t, err)
		assert.Len(t, listCallbacksResponse.Callbacks, 0)
	})
}
