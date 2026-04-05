package core

import "time"

type LogEntry struct {
	Level     LogLevel
	Timestamp time.Time
	Message   string
}
