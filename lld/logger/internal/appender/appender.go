package appender

import (
	"lld-logger/internal/core"
)

type Appender interface {
	Append(entry core.LogEntry)
}
