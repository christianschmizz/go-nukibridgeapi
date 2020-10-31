package bridge

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nukibridge "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func createUnlockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <deviceType> <deviceID>",
		Short: "Unlock",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			nukiID, err := resolveNukiIDFromArgs(args)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to resolve nukiID from args")
			}

			conn, err := nukibridge.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
			}

			lock, err := conn.Unlock(*nukiID)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to unlock")
			}
			fmt.Printf("%+v\n", lock)
		},
	}

	return cmd
}
