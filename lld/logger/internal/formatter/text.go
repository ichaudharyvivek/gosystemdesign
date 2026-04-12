package formatter

import (
	"fmt"
	"lld-logger/internal/model"
	"strings"
	"time"
)

type TextFormatter struct{}

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

func (f *TextFormatter) Format(entry model.Record) []byte {
	var b strings.Builder

	b.WriteString(entry.Timestamp.Format(time.RFC3339))
	b.WriteString(" [")
	b.WriteString(entry.Level.String())
	b.WriteString("] ")
	b.WriteString(entry.Message)

	if len(entry.Fields) > 0 {
		b.WriteString(" | ")
		sep := ""
		for k, v := range entry.Fields {
			b.WriteString(sep)
			fmt.Fprintf(&b, "'%s': '%+v'", k, v)
			sep = ", "
		}
	}

	b.WriteString("\n")

	return []byte(b.String())
}
