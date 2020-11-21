package watch

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	dBusPath          dbus.ObjectPath = "/nuki/bridge"
	dBusInterfaceName string          = "nuki.bridge.Event"
)

func forwarder(ctx context.Context, wg *sync.WaitGroup, ifname, nukiBridgeHost, nukiBridgeToken string) {
	defer wg.Done()

	// Lookup local interface address required for creating callbacks
	localhostAddress, err := lookupNetworkInterfaceIPv4Address(ifname)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to lookup interface")
	}
	log.Debug().Str("ifname", ifname).Str("address", localhostAddress).Msg("resolved interface address")

	// Connect to Nuki bridge
	conn, err := bridgeapi.ConnectWithToken(nukiBridgeHost, nukiBridgeToken)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
	}

	// Connect to dbus' session bus
	dbusConn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to session bus:")
	}
	defer func() {
		if err := dbusConn.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close dbus connection")
		}
	}()

	rand.Seed(time.Now().UnixNano())
	callbackPath := fmt.Sprintf("/callback/%d", rand.Intn(999999-100000)+100000)
	callbackURL := fmt.Sprintf("http://%s:8080%s", localhostAddress, callbackPath)

	// Setup callbacks at the bridge
	callbackLogger := log.With().Str("callback_url", callbackURL).Logger()
	registerCallback(conn, callbackURL, callbackLogger)
	defer unregisterCallback(conn, callbackURL, callbackLogger)

	mux := http.NewServeMux()
	mux.HandleFunc(callbackPath, createCallbackRequestHandler(dbusConn))

	wg.Add(1)
	go func() {
		wg.Done()

		httpAddr := localhostAddress + ":8080"
		log.Info().Str("addr", httpAddr).Msg("http server starts listening")
		if err = http.ListenAndServe(httpAddr, mux); err != nil {
			log.Fatal().Err(err).Msg("http server reported an error")
		}
	}()

	<-ctx.Done()
}

func registerCallback(conn *bridgeapi.Connection, callbackURL string, logger zerolog.Logger) {
	resp, err := conn.AddCallback(callbackURL)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to register callback")
	}
	if !resp.Success {
		logger.Fatal().Err(err).Msg("failed to add callback")
	}
	log.Debug().Str("callback_url", callbackURL).Msg("successfully registered callback")
}

func unregisterCallback(conn *bridgeapi.Connection, callbackURL string, callbackLogger zerolog.Logger) {
	log.Debug().Str("callback_url", callbackURL).Msg("unregistering callback")

	callbacks, err := conn.ListCallbacks()
	if err != nil {
		callbackLogger.Fatal().Err(err).Msg("failed to fetch callbacks")
	}

	for _, callback := range callbacks.Callbacks {
		if callback.URL == callbackURL {
			removeResp, err := conn.RemoveCallback(callback.ID)
			if err != nil {
				callbackLogger.Fatal().Err(err).Msg("failed to remove callback")
			}
			if !removeResp.Success {
				callbackLogger.Fatal().Err(err).Msg("removing the callback failed")
			}
			return
		}
	}

	callbackLogger.Warn().Msg("no matching callback found to remove")
}
