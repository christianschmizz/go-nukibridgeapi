package bridge

import (
	"fmt"
	"strconv"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/cmd/bridge/watch"
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

// resolveNukiIDFromArgs assembles an ID from the first two args
func resolveNukiIDFromArgs(args []string) (*nuki.ID, error) {
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
