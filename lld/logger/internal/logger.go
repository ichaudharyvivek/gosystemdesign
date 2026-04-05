package logger

import (
	"lld-logger/internal/appender"
	"lld-logger/internal/core"
	"time"
)

type Logger struct {
	level     core.LogLevel
	appenders []appender.Appender
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) SetLevel(level core.LogLevel) {
	l.level = level
}

func (l *Logger) SetAppenders(appenders []appender.Appender) {
	l.appenders = appenders
}

func (l *Logger) Debug(message string) {
	l.log(core.DebugLevel, message)
}

func (l *Logger) Info(message string) {
	l.log(core.InfoLevel, message)
}

func (l *Logger) Warn(message string) {
	l.log(core.WarnLevel, message)
}

func (l *Logger) Error(message string) {
	l.log(core.ErrorLevel, message)
}

func (l *Logger) Fatal(message string) {
	l.log(core.FatalLevel, message)
}

func (l *Logger) log(level core.LogLevel, message string) {
	if level >= l.level {
		entry := core.LogEntry{
			Level:     level,
			Message:   message,
			Timestamp: time.Now(),
		}

		for _, appender := range l.appenders {
			appender.Append(entry)
		}
	}
}
