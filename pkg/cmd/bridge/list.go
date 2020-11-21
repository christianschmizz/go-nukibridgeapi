package bridge

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

func createListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List paired devices",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := bridgeapi.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
			if err != nil {
				log.Fatal().Err(err).Msg("")
			}

			devices, err := conn.ListPairedDevices()
			if err != nil {
				log.Fatal().Err(err).Msg("")
			}

			printDeviceList(os.Stdout, devices)
		},
	}
	return cmd
}

func printDeviceList(writer io.Writer, devices bridgeapi.ListPairedDevicesResponse) {
	w := tabwriter.NewWriter(writer, 3, 0, 1, ' ', 0)
	defer w.Flush()

	_, _ = fmt.Fprintln(w, "Type\tID\tName\tBattery\tFirmware Version")
	for _, d := range devices {
		_, _ = fmt.Fprintf(w, "%d\t%d\t%s\t%d%%\t%s\n", d.Type, d.ID, d.Name, d.LastKnownState.BatteryChargeState, d.FirmwareVersion)
	}
}
