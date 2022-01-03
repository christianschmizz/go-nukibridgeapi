package bridge

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func createUnlockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <deviceType> <deviceID>",
		Short: "Unlock",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			deviceID := mustResolveDeviceIDFromArgs(args)
			lock, err := mustConnect(nil).Unlock(*deviceID)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to unlock")
			}
			fmt.Printf("%+v\n", lock)
		},
	}

	return cmd
}
