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

func (a *ConsoleAppender) Append(data []byte) {
	a.out.Write(data)
}
