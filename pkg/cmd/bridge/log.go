package bridge

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func createLogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log",
		Short: "Log entries",
		Run: func(cmd *cobra.Command, args []string) {
			offset, _ := cmd.Flags().GetInt("offset")
			count, _ := cmd.Flags().GetInt("count")
			logLines, err := mustConnect(nil).Log(offset, count)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to retrieve log entries from Nuki bridge")
			}

			printLog(os.Stdout, logLines)
		},
	}

	cmd.Flags().Int("offset", 0, "Offset")
	cmd.Flags().Int("count", 100, "Count")

	return cmd
}

func printLog(writer io.Writer, logLines bridgeapi.Log) {
	w := tabwriter.NewWriter(writer, 3, 0, 1, ' ', 0)
	defer w.Flush()

	_, _ = fmt.Fprintln(w, "Time\tID\tType\tHandle\tMac address")
	for _, l := range logLines {
		_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", l.Timestamp.UTC(), l.ID, l.Type, l.BleHandle, l.MacAddr)
	}
}
