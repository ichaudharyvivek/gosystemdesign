package appender

import (
	"io"
	"lld-logger/internal/formatter"
	"lld-logger/internal/model"
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

func (a *ConsoleAppender) Append(record model.Record) error {
	data := a.formatter.Format(record)
	_, err := a.out.Write(data)
	return err
}

func (a *ConsoleAppender) Close() error {
	return nil
}
