package bridge

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func createInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Retrieve bridge info",
		Run: func(cmd *cobra.Command, args []string) {
			info, err := mustConnect(nil).Info()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to retrieve info from Nuki bridge")
			}

			printVersions(os.Stdout, info)
		},
	}

	return cmd
}

func printVersions(writer io.Writer, info *bridgeapi.InfoResponse) {
	if PrintJSON {
		if data, err := json.MarshalIndent(info, "", "   "); err == nil {
			fmt.Println(string(data))
		}
	} else {
		if bridgeapi.HasSupportForEncryptedTokens(info.BridgeType, info.Versions.FirmwareVersion) {
			fmt.Printf("This bridge supports encrypted tokens.\n\n")
		}

		printScanResults(os.Stdout, info.ScanResults)
	}
}

func printScanResults(writer io.Writer, results []bridgeapi.ScanResult) {
	w := tabwriter.NewWriter(writer, 3, 0, 1, ' ', 0)
	_, _ = fmt.Fprintln(w, "ID\tType\tName\tRSSI\tPaired")
	for _, result := range results {
		_, _ = fmt.Fprintf(w, "%d\t%d\t%s\t%d\t%t\n", result.ID, result.Type, result.Name, result.Rssi, result.Paired)
	}
	_ = w.Flush()
}
