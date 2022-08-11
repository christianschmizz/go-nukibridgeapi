package root

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/christianschmizz/go-nukibridgeapi/pkg/cmd/bridge"
	"github.com/christianschmizz/go-nukibridgeapi/pkg/cmd/discover"
)

var (
	cfgFile string
	noViper bool
)

// createCommand creates the root command
func createCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nukibridgectl <command> <subcommand> [flags]",
		Short: "Nuki Bridge CLI",
		Long:  `Work seamlessly with your Nuki Bridge from the command line.`,
	}

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nukibridgectl.yaml)")
	cmd.PersistentFlags().BoolVar(&noViper, "noViper", false, "disable the use of viper")

	cmd.AddCommand(bridge.CreateCommand())
	cmd.AddCommand(discover.CreateCommand())

	return cmd
}

// Execute command
func Execute() {
	rootCmd := createCommand()

	// Initialize viper when initializing any cobra commands
	cobra.OnInitialize(initViper)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("failed to execute command")
	}

	os.Exit(0)
}

func initViper() {
	if noViper {
		return
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal().Err(err).Msg("no home directory found")
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
		log.Warn().Err(err).Msg("could not read config")
	}

	log.Info().Str("config_file", viper.ConfigFileUsed()).Msg("using config file")
}
