package bridge

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
	_ "github.com/godbus/dbus/v5"
	"github.com/google/go-cmp/cmp"
)

const (
	dBusPath          dbus.ObjectPath = "/nuki/bridge"
	dBusInterfaceName string          = "nuki.bridge.Event"
)

var (
	lastState map[nuki.ID]CallbackData
)

func init() {
	lastState = make(map[nuki.ID]CallbackData)
}

type CallbackData struct {
	ID                    int                  `json:"nukiId"`
	Type                  nuki.DeviceType      `json:"deviceType"`
	Mode                  nuki.LockMode        `json:"mode"`
	State                 nuki.LockState       `json:"state"`
	StateName             string               `json:"stateName"`
	BatteryCritical       bool                 `json:"batteryCritical"`
	KeypadBatteryCritical bool                 `json:"keypadBatteryCritical"`
	DoorsensorState       nuki.DoorsensorState `json:"doorsensorState,omitempty"`
	DoorsensorStateName   string               `json:"doorsensorStateName,omitempty"`
	RingactionTimestamp   time.Time            `json:"ringactionTimestamp,omitempty" dbus:"-"` // Encoding of timestamps leads to an error '"dbus: connection closed by user"'
	RingactionState       bool                 `json:"ringactionState,omitempty"`
}

// NukiID assembles the ID from a result
func (d *CallbackData) NukiID() *nuki.ID {
	return &nuki.ID{
		DeviceID:   d.ID,
		DeviceType: d.Type,
	}
}

func createWatchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch <localInterfaceName>",
		Short: "Watch bridge for changes and emits them to DBus",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)

			ifname := args[0]

			conn, err := bridgeapi.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
			}

			host, err := lookupNetworkInterfaceIPv4Address(ifname)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to lookup interface")
			}

			dbusConn, err := dbus.SessionBus()
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to connect to session bus:")
			}
			defer dbusConn.Close()

			rand.Seed(time.Now().UnixNano())
			callbackPath := fmt.Sprintf("/callback/%d", rand.Intn(999999-100000)+100000)
			callbackURL := fmt.Sprintf("http://%s:8080%s", host, callbackPath)

			{
				callbackLogger := log.With().Str("callback_url", callbackURL).Logger()

				{
					resp, err := conn.AddCallback(callbackURL)
					if err != nil {
						callbackLogger.Fatal().Err(err).Msg("failed to register callback")
					}
					if !resp.Success {
						callbackLogger.Fatal().Err(err).Msg("adding callback was not successful")
					}
					log.Debug().Str("callback_url", callbackURL).Msg("successfully registered callback")
				}

				defer func() {
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
				}()
			}

			http.HandleFunc(callbackPath, createCallbackHandler(dbusConn))

			httpAddr := host + ":8080"

			go func() {
				if err = http.ListenAndServe(httpAddr, nil); err != nil {
					log.Fatal().Err(err).Msg("http server reported an error")
				}
			}()

			log.Info().Str("addr", httpAddr).Msg("http server is listening")

			// Wait for signal to shutdown
			<-c
		},
	}

	return cmd
}

func createCallbackHandler(conn *dbus.Conn) func(writer http.ResponseWriter, req *http.Request) {
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

func isValidIPv4(ipAddress string) bool {
	testInput := net.ParseIP(ipAddress)
	return testInput.To4() != nil
}

// lookupNetworkInterfaceIPv4Address return the first valid IPv4 address to the given network interface
func lookupNetworkInterfaceIPv4Address(name string) (string, error) {
	byNameInterface, err := net.InterfaceByName(name)
	if err != nil {
		return "", fmt.Errorf("network interface not found: %s", name)
	}

	addresses, err := byNameInterface.Addrs()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve addresses for interface: %s", name)
	}

	for _, addr := range addresses {
		tokens := strings.Split(addr.String(), "/")
		if isValidIPv4(tokens[0]) {
			return tokens[0], nil
		}
	}

	return "", fmt.Errorf("no addresses found for interface: %s", name)
}
