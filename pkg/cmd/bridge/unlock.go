package bridge

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
	nukibridge "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func createUnlockCommand() (*cobra.Command) {
	cmd := &cobra.Command{
		Use:   "unlock <deviceType> <deviceID>",
		Short: "Unlock",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			deviceType, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal().Err(err).Msg("invalid device's type")
			}

			deviceID, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal().Err(err).Msg("invalid device's ID")
			}

			nukiID := nuki.NukiID{deviceID, nuki.DeviceType(deviceType)}

			conn, err := nukibridge.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
			}

			lock, err := conn.Unlock(nukiID)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to unlock")
			}
			fmt.Printf("%+v\n", lock)
		},
	}

	return cmd
}