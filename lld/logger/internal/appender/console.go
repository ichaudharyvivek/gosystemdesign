package appender

import (
	"io"
)

type ConsoleAppender struct {
	out io.Writer
}

func NewConsoleAppender(out io.Writer) *ConsoleAppender {
	return &ConsoleAppender{
		out: out,
	}
}

func (a *ConsoleAppender) Append(data []byte) error {
	_, err := a.out.Write(data)
	return err
}

func (a *ConsoleAppender) Close() error {
	return nil
}
