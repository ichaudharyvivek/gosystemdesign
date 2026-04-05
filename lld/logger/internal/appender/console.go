package appender

import (
	"io"
	"lld-logger/internal/core"
	"lld-logger/internal/formatter"
)

type ConsoleAppender struct {
	out       io.Writer
	formatter formatter.Formatter
}

func NewConsoleAppender(out io.Writer, formatter formatter.Formatter) *ConsoleAppender {
	return &ConsoleAppender{
		out:       out,
		formatter: formatter,
	}
}

func (a *ConsoleAppender) Append(entry core.LogEntry) {
	data := a.formatter.Format(entry)
	a.out.Write(data)
}
