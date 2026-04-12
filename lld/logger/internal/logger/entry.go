package logger

import (
	"fmt"
	"lld-logger/internal/model"
	"os"
	"time"
)

type Entry struct {
	logger    *Logger
	level     model.Level
	timestamp time.Time
	message   string
	fields    map[string]any
}

func (e *Entry) WithFields(fields map[string]any) *Entry {
	if e.fields == nil {
		e.fields = make(map[string]any)
	}

	for k, v := range fields {
		e.fields[k] = v
	}

	return e
}

func (e *Entry) Msg(message string) {
	if e.level < e.logger.level {
		return
	}

	e.message = message
	e.timestamp = time.Now()
	record := model.Record{
		Level:     e.level,
		Fields:    e.fields,
		Message:   e.message,
		Timestamp: e.timestamp,
	}

	for _, appender := range e.logger.appenders {
		if err := appender.Append(record); err != nil {
			fmt.Fprintln(os.Stderr, "log write error:", err)
		}
	}

	if e.level == model.FatalLevel {
		os.Exit(1)
	}
}
