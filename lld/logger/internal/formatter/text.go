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

func (f *TextFormatter) Format(record model.Record) []byte {
	var b strings.Builder

	b.WriteString(record.Timestamp.Format(time.RFC3339))
	b.WriteString(" [")
	b.WriteString(record.Level.String())
	b.WriteString("] ")
	b.WriteString(record.Message)

	if len(record.Fields) > 0 {
		b.WriteString(" | ")
		sep := ""
		for k, v := range record.Fields {
			b.WriteString(sep)
			fmt.Fprintf(&b, "'%s': '%+v'", k, v)
			sep = ", "
		}
	}

	b.WriteString("\n")

	return []byte(b.String())
}
