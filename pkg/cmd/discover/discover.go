package discover

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/nuki/bridgeapi"
)

// CreateCommand create the "discover" command
func CreateCommand() *cobra.Command {
	discoverCmd := &cobra.Command{
		Use:   "discover",
		Short: "Discover bridges",
		Run: func(cmd *cobra.Command, args []string) {
			discovery, err := bridgeapi.Discover()
			if err != nil {
				log.Fatal().Err(err).Msg("discovery failed")
			}
			log.Info().Msgf("found %d bridges", len(discovery.Bridges))

			printBridges(os.Stdout, discovery.Bridges)
		},
	}

	return discoverCmd
}

func printBridges(writer io.Writer, bridges []bridgeapi.BridgeInfo) {
	w := tabwriter.NewWriter(writer, 3, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tIP\tPort\tUpdated")
	for _, bridge := range bridges {
		fmt.Fprintf(w, "%d\t%s\t%d\t%s\n", bridge.BridgeID, bridge.IP, bridge.Port, bridge.DateUpdated)
	}
	w.Flush()
}
