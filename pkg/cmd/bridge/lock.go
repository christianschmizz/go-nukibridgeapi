package bridge

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

func createLockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock <deviceType> <deviceID>",
		Short: "Lock",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			nukiID, err := resolveNukiIDFromArgs(args)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to resolve nukiID from args")
			}

			conn, err := bridgeapi.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
			}

			lock, err := conn.Lock(*nukiID)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to lock")
			}
			fmt.Printf("%+v\n", lock)
		},
	}

	return cmd
}
