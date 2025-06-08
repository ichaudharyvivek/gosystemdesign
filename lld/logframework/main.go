package main

import (
	c "lld/logframework/core"
	l "lld/logframework/logger"
)

var (
	cfg = &l.LogConfig{
		Level: l.INFO,
		Appenders: []l.Appender{
			&c.ConsoleAppender{},
		},
	}
)

func main() {
	log := l.GetLogger()
	log.SetConfiguration(cfg)

	log.Trace("This is trace. It will not be logged")
	log.Info("This is info, and is going to be logged")

	const CLEVEL l.LogLevel = "CUSTOM LEVEL"
	log.WithLevel(CLEVEL, "This is error, this too is going to be logged")
}
