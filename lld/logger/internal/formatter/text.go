package formatter

import (
	"lld-logger/internal/core"
	"strings"
	"time"
)

type TextFormatter struct{}

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

func (tf *TextFormatter) Format(entry core.LogEntry) []byte {
	var b strings.Builder

	b.WriteString(entry.Timestamp.Format(time.RFC3339))
	b.WriteString(" [")
	b.WriteString(entry.Level.String())
	b.WriteString("] ")
	b.WriteString(entry.Message)
	b.WriteString("\n")

	return []byte(b.String())
}
