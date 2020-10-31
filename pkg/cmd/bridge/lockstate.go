package bridge

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	api "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func createLockStateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lockState <deviceType> <deviceID>",
		Short: "Pull the lockState directly of the device",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			nukiID, err := resolveNukiIDFromArgs(args)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to resolve nukiID from args")
			}

			conn, err := api.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
			}
			state, err := conn.LockState(*nukiID)
			if err != nil {
				log.Fatal().Err(err).Msg("")
			}
			fmt.Printf("%+v\n", state)
		},
	}

	return cmd
}
