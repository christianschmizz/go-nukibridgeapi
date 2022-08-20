package bridge

import (
	"fmt"
	"strconv"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/cmd/bridge/watch"
	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
)

var (
	// Host is the address and port of the bridge (e.g. "192.168.0.1:8080")
	Host string

	// Token is the Auth token required for accessing the bridge's API
	Token string

	// UseEncryptedTokens enables the use of encrypted tokens (only supported by Bridge 1.0: ≥1.22.1 and Bridge 2.0: ≥2.14.0)
	UseEncryptedTokens bool
)

// CreateCommand creates the "bridge" command group
func CreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge",
		Short: "Bridge commands",
	}

	cmd.PersistentFlags().StringVarP(&Host, "host", "b", "", "Host and port")
	if err := viper.BindPFlag("host", cmd.PersistentFlags().Lookup("host")); err != nil {
		log.Fatal().Err(err).Msg("unable to bind flag")
	}

	cmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "Bridge token for auth")
	if err := viper.BindPFlag("token", cmd.PersistentFlags().Lookup("token")); err != nil {
		log.Fatal().Err(err).Msg("unable to bind flag")
	}

	cmd.PersistentFlags().BoolVar(&UseEncryptedTokens, "encrypt", false, "Use encrypted tokens")
	if err := viper.BindPFlag("encrypt", cmd.PersistentFlags().Lookup("encrypt")); err != nil {
		log.Fatal().Err(err).Msg("unable to bind flag")
	}

	// Child commands
	cmd.AddCommand(createInfoCommand())
	cmd.AddCommand(createListCommand())
	cmd.AddCommand(createLogCommand())

	// Require nukiID
	cmd.AddCommand(createLockActionCommand())
	cmd.AddCommand(createLockStateCommand())
	cmd.AddCommand(createLockCommand())
	cmd.AddCommand(createUnlockCommand())
	// cmd.AddCommand(createUnpairCommand())
	cmd.AddCommand(watch.CreateWatchCommand())

	callbacks := &cobra.Command{Use: "callbacks", Short: "Callback management"}
	callbacks.AddCommand(createCallbackListCommand())
	callbacks.AddCommand(createCallbackAddCommand())
	callbacks.AddCommand(createCallbackRemoveCommand())

	cmd.AddCommand(callbacks)

	return cmd
}

// resolveDeviceIDFromArgs assembles an ID from the first two args
func resolveDeviceIDFromArgs(args []string) (*nuki.ID, error) {
	deviceType, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, fmt.Errorf("invalid device's type: %v", args[0])
	}

	deviceID, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, fmt.Errorf("invalid device's ID: %v", args[1])
	}

	return &nuki.ID{DeviceID: deviceID, DeviceType: nuki.DeviceType(deviceType)}, nil
}

func mustResolveDeviceIDFromArgs(args []string) *nuki.ID {
	nukiID, err := resolveDeviceIDFromArgs(args)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to resolve nukiID from args")
	}
	return nukiID
}

func mustConnect(v *viper.Viper) *bridgeapi.Connection {
	if v == nil {
		v = viper.GetViper()
	}
	options := make([]func(*bridgeapi.Connection), 0)
	if v.GetBool("encrypt") {
		options = append(options, bridgeapi.UseEncryptedToken())
	}
	conn, err := bridgeapi.ConnectWithToken(v.GetString("host"), v.GetString("token"), options...)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
	}
	return conn
}
