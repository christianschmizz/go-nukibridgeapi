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

func createCallbackListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List callbacks",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := mustConnect(nil).ListCallbacks()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to retrieve list of callbacks from Nuki bridge")
			}

			printCallbacks(os.Stdout, resp.Callbacks)
		},
	}

	return cmd
}

func printCallbacks(writer io.Writer, callbacks []bridgeapi.Callback) {
	if len(callbacks) > 0 {
		w := tabwriter.NewWriter(writer, 3, 0, 1, ' ', 0)
		defer w.Flush()

		_, _ = fmt.Fprintln(w, "ID\tURL")
		for _, callback := range callbacks {
			_, _ = fmt.Fprintf(w, "%d\t%s\n", callback.ID, callback.URL)
		}
	} else {
		fmt.Println("no callbacks found")
	}

}
