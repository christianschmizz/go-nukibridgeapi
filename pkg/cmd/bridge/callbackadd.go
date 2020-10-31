package bridge

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	api "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func createCallbackAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <url>",
		Short: "List callbacks",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]

			conn, err := api.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
			}

			resp, err := conn.AddCallback(url)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to retrieve add callback to Nuki bridge")
			}
			fmt.Printf("$%v", resp)
		},
	}

	return cmd
}
