package main

import (
	logger "lld-logger/internal"
	appender "lld-logger/internal/appender"
	core "lld-logger/internal/core"
	"lld-logger/internal/formatter"
	"os"
)

func main() {
	log := logger.New()
	log.SetLevel(core.InfoLevel)
	log.SetAppenders([]appender.Appender{
		appender.NewConsoleAppender(os.Stdout, formatter.NewTextFormatter()),
	})

	// Writting logs
	log.Debug("This is a debug message")
	log.Info("This is an info messsage")
	log.Warn("This is a warning message")
	log.Error("This is an error message")
	log.Fatal("This is a fatal message")
}
