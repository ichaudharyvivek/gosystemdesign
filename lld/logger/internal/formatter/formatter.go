package formatter

import "lld-logger/internal/core"

type Formatter interface {
	Format(entry core.LogEntry) []byte
}
