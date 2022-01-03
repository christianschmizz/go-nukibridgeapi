package bridge

import (
	"fmt"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func createLockActionCommand() *cobra.Command {
	var NoWait bool
	cmd := &cobra.Command{
		Use:   "lockAction <deviceType> <deviceID> <action>",
		Short: "Execute a lockAction",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			deviceID := mustResolveDeviceIDFromArgs(args)

			actionName := args[2]
			action, err := nuki.LockActionFromString(actionName, deviceID.DeviceType)
			if err != nil {
				log.Fatal().Err(err).Str("action", actionName).Msg("invalid action")
			}

			state, err := mustConnect(nil).LockAction(*deviceID, action)
			if err != nil {
				log.Error().Err(err).Msg("failed to issue request")
			}
			fmt.Printf("%t\n", state.Success)
		},
	}
	cmd.Flags().BoolVarP(&NoWait, "nowait", "n", true, "Don't wait until the command finished")

	return cmd
}
