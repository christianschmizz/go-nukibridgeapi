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

func createInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Retrieve bridge info",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := bridgeapi.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
			}

			info, err := conn.Info()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to retrieve info from Nuki bridge")
			}

			printScanResults(os.Stdout, info.ScanResults)
		},
	}

	return cmd
}

func printScanResults(writer io.Writer, results []bridgeapi.ScanResult) {
	w := tabwriter.NewWriter(writer, 3, 0, 1, ' ', 0)
	_, _ = fmt.Fprintln(w, "ID\tType\tName\tRSSI\tPaired")
	for _, result := range results {
		_, _ = fmt.Fprintf(w, "%d\t%d\t%s\t%d\t%t\n", result.ID, result.Type, result.Name, result.Rssi, result.Paired)
	}
	w.Flush()
}
