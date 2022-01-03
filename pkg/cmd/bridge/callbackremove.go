package bridge

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func createCallbackRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove <id>",
		Short: "Remove callback",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			callbackID, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal().Err(err).Msg("invalid callbackID")
			}

			resp, err := mustConnect(nil).RemoveCallback(callbackID)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to remove callback from Nuki bridge")
			}
			fmt.Printf("$%v", resp)
		},
	}

	return cmd
}
