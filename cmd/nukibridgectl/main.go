package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/christianschmizz/go-nukibridgeapi/internal/build"
	"github.com/christianschmizz/go-nukibridgeapi/pkg/cmd/root"
)

func main() {
	hasDebug := os.Getenv("DEBUG") != ""

	// Default level for this example is info, unless debug is active
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	if hasDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	log.Info().
		Str("build_date", build.Date).
		Str("build_version", build.Version).
		Bool("debug_mode", hasDebug).
		Msg("booting")

	root.Execute()
}
