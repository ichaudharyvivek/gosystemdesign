package core

import (
	"strconv"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota + 1
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func (lvl LogLevel) String() string {
	switch lvl {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "unknown (" + strconv.Itoa(int(lvl)) + ")"
	}
}
