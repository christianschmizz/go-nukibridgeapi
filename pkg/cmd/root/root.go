package root

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/cmd/bridge"
	"github.com/christianschmizz/go-nukibridgeapi/pkg/cmd/discover"
)

var (
	cfgFile  string
	useViper bool
)

func CreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nukibridgectl <command> <subcommand> [flags]",
		Short: "Nuki Bridge CLI",
		Long:  `Work seamlessly with your Nuki Bridge from the command line.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nukibridgectl.yaml)")

	cmd.AddCommand(bridge.CreateCommand())
	cmd.AddCommand(discover.CreateCommand())

	return cmd
}

func init() {
	// Run before every command
	cobra.OnInitialize(initViper)
}

func initViper() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			panic(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		for _, p := range []string{"/etc/nukibridgectl/", home, "."} {
			viper.AddConfigPath(p) // path to look for the config file in
		}

		viper.SetConfigName(".nukibridgectl")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Fprintf(os.Stderr, "Config file not found; ignore error if desired")
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stderr, "Config file was found but another error was produced")
			os.Exit(1)
		}
	}

	log.Info().Str("config_file", viper.ConfigFileUsed()).Msg("using config file")
}
