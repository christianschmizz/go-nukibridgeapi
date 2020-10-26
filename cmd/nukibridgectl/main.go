package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/christianschmizz/go-nukibridgeapi/internal/build"
	"github.com/christianschmizz/go-nukibridgeapi/pkg/cmd/root"
)

func main() {
	buildDate := build.Date
	buildVersion := build.Version

	hasDebug := os.Getenv("DEBUG") != ""

	// Default level for this example is info, unless debug is active
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if hasDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	log.Info().Str("buildDate", buildDate).Str("buildVersion", buildVersion).Bool("debugMode", hasDebug).Msg("")

	if err := root.CreateCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
