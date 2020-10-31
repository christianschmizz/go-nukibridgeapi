package bridge

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	api "github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridge"
)

func createCallbackListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List callbacks",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := api.ConnectWithToken(viper.GetString("host"), viper.GetString("token"))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to connect to Nuki bridge")
			}

			resp, err := conn.ListCallbacks()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to retrieve list of callbacks from Nuki bridge")
			}

			printCallbacks(os.Stdout, resp.Callbacks)
		},
	}

	return cmd
}

func printCallbacks(writer io.Writer, callbacks []api.Callback) {
	if len(callbacks) > 0 {
		w := tabwriter.NewWriter(writer, 3, 0, 1, ' ', 0)
		defer w.Flush()

		fmt.Fprintln(w, "ID\tURL")
		for _, callback := range callbacks {
			fmt.Fprintf(w, "%d\t%s\n", callback.ID, callback.URL)
		}
	} else {
		fmt.Println("no callbacks found")
	}

}
