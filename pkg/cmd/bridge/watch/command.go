package watch

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CreateWatchCommand creates the cobra command for watching the bridge
func CreateWatchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch <localInterfaceName>",
		Short: "Watch bridge for changes and emits them to DBus",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ifname := args[0]

			var wg sync.WaitGroup
			ctx, cancel := context.WithCancel(context.Background())

			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)

			wg.Add(1)
			go forwarder(ctx, &wg, ifname, viper.GetString("host"), viper.GetString("token"))

			// Wait for signal to shutdown
			<-c
			cancel()
			wg.Wait()
		},
	}

	return cmd
}
