package bridge

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
	api "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func createLockActionCommand() *cobra.Command {
	var NoWait bool
	cmd := &cobra.Command{
		Use:   "lockAction <deviceType> <deviceID> <action>",
		Short: "Execute a lockAction",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			nukiID, err := resolveNukiIDFromArgs(args)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to resolve nukiID from args")
			}

			actionName := args[1]
			action, err := nuki.SmartLockActionFromString(actionName)
			if err != nil {
				log.Fatal().Err(err).Str("action", actionName).Msg("invalid action")
			}

			conn, err := api.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
			}

			state, err := conn.LockAction(*nukiID, action)
			if err != nil {
				log.Fatal().Err(err).Msg("")
			}
			fmt.Printf("%+v\n", state)
		},
	}
	cmd.Flags().BoolVarP(&NoWait, "nowait", "n", true, "Wait until command finished")

	return cmd
}
