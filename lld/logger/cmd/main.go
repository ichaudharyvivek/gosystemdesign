package main

import (
	logger "lld-logger/internal/logger"
)

func main() {
	log := logger.New()

	// Writting logs
	log.Debug().Msg("This is a debug message")
	log.Info().Fields(map[string]any{"type": "info", "stackName": "ANT_UAT"}).Msg("This is an info messsage")
	log.Warn().Msg("This is a warning message")
	log.Error().Msg("This is an error message")
	log.Fatal().Msg("This is a fatal message")
}
