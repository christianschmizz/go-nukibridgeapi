package bridge

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func createCallbackAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <url>",
		Short: "List callbacks",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]

			resp, err := mustConnect(nil).AddCallback(url)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to retrieve add callback to Nuki bridge")
			}
			fmt.Printf("$%v", resp)
		},
	}

	return cmd
}
