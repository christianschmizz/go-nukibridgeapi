package bridge

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Host  string
	Token string
)

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

	callbacks := &cobra.Command{Use: "callbacks", Short: "Callback management"}
	callbacks.AddCommand(createCallbackListCommand())
	callbacks.AddCommand(createCallbackAddCommand())
	callbacks.AddCommand(createCallbackRemoveCommand())

	cmd.AddCommand(callbacks)

	return cmd
}
