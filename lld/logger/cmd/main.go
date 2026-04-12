package main

import (
	"lld-logger/internal/logger"

	"golang.org/x/sync/errgroup"
)

func main() {
	TestSimple()
	// TestConcurrent()
}

func TestSimple() {
	log := logger.New()
	defer log.Close()

	log.Debug().Msg("This is a debug message")
	log.Info().WithFields(map[string]any{"type": "info", "stackName": "ANT_UAT"}).Msg("This is an info messsage")
	log.Warn().Msg("This is a warning message")
	log.Error().Msg("This is an error message")
	log.Fatal().Msg("This is a fatal message")
}

func TestConcurrent() {
	// Run with: go run -race main.go
	log := logger.New()
	defer log.Close()

	var g errgroup.Group
	for i := 0; i < 100; i++ {
		i := i
		g.Go(func() error {
			log.Info().WithFields(map[string]any{
				"id": i,
			}).Msg("concurrent log")
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		panic(err)
	}
}
