package bridge

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nukibridge "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func createListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List paired devices",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := nukibridge.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
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

func printDeviceList(writer io.Writer, devices nukibridge.ListPairedDevicesResponse) {
	w := tabwriter.NewWriter(writer, 3, 0, 1, ' ', 0)
	defer w.Flush()

	_, _ = fmt.Fprintln(w, "ID\tType\tName\tBattery")
	for _, d := range devices {
		_, _ = fmt.Fprintf(w, "%d\t%d\t%s\t%d%%\n", d.ID, d.Type, d.Name, d.LastKnownState.BatteryChargeState)
	}
}
