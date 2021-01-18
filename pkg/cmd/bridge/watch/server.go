package watch

import (
	"encoding/json"
	"net/http"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
	"github.com/godbus/dbus/v5"
	"github.com/google/go-cmp/cmp"
	"github.com/rs/zerolog/log"
)

var (
	lastState map[nuki.ID]CallbackData
)

func init() {
	lastState = make(map[nuki.ID]CallbackData)
}

func createCallbackRequestHandler(conn *dbus.Conn) func(writer http.ResponseWriter, req *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		var data CallbackData

		// Read and decode callback's data
		if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
			log.Debug().Err(err).Str("remote_server", req.RemoteAddr).Msg("failed to decode data")
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		defer req.Body.Close()

		// Diff states
		id := *data.NukiID()
		if oldData, ok := lastState[id]; ok {
			if diff := cmp.Diff(data, oldData); diff != "" {
				log.Debug().Msg(diff)
			}
		}
		lastState[id] = data

		// Broadcast callback data
		if err := conn.Emit(dBusPath, dBusInterfaceName, data); err != nil {
			log.Error().Err(err).Msg("failed to emit signal")
		}
	}
}
