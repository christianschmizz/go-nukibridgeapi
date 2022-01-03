package bridge

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func createLockStateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lockState <deviceType> <deviceID>",
		Short: "Pull the lockState directly of the device",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			deviceID := mustResolveDeviceIDFromArgs(args)
			state, err := mustConnect(nil).LockState(*deviceID)
			if err != nil {
				log.Fatal().Err(err).Msg("")
			}
			fmt.Printf("%+v\n", state)
		},
	}

	return cmd
}
