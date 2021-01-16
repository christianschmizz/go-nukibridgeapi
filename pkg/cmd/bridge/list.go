package bridge

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

var showLastKnownState bool

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
	cmd.Flags().BoolVar(&showLastKnownState, "show-last-known-state", false, "Show details about last known state of the devices (if available)")
	return cmd
}

func printDeviceList(writer io.Writer, devices bridgeapi.ListPairedDevicesResponse) {
	w := tabwriter.NewWriter(writer, 3, 0, 1, ' ', 0)
	defer w.Flush()

	headers := []string{"Type", "ID", "Name", "Battery", "Firmware Version"}
	if showLastKnownState {
		headers = append(headers, "State", "BCrit", "BChrg", "BChrgSt", "KPadCrit", "DoorSt", "RingActSt", "RingActTimestamp", "Timestamp")
	}
	_, _ = io.WriteString(w, strings.Join(headers, "\t")+"\n")

	for _, d := range devices {
		_, _ = fmt.Fprintf(w, "%d\t%d\t%s\t%d%%\t%s", d.Type, d.ID, d.Name, d.LastKnownState.BatteryChargeState, d.FirmwareVersion)
		if showLastKnownState {
			_, _ = fmt.Fprintf(w, "\t%s\t%t\t%t\t%d\t%t\t%d\t%t\t%s\t%s",
				d.LastKnownState.StateName,
				d.LastKnownState.BatteryCritical,
				d.LastKnownState.BatteryCharging,
				d.LastKnownState.BatteryChargeState,
				d.LastKnownState.KeypadBatteryCritical,
				d.LastKnownState.DoorsensorState,
				d.LastKnownState.RingactionState,
				d.LastKnownState.RingactionTimestamp,
				d.LastKnownState.Timestamp,
			)
		}
		_, _ = io.WriteString(w, "\n")
	}
}
