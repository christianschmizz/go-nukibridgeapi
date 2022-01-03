package bridge

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func createLockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock <deviceType> <deviceID>",
		Short: "Lock",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			deviceID := mustResolveDeviceIDFromArgs(args)
			lock, err := mustConnect(nil).Lock(*deviceID)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to lock")
			}
			fmt.Printf("%+v\n", lock)
		},
	}

	return cmd
}
