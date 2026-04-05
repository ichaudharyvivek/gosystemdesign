package core

import (
	"fmt"
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

func ParseLevel(s string) (LogLevel, error) {
	switch s {
	case "DEBUG":
		return DebugLevel, nil
	case "INFO":
		return InfoLevel, nil
	case "WARN":
		return WarnLevel, nil
	case "ERROR":
		return ErrorLevel, nil
	case "FATAL":
		return FatalLevel, nil
	default:
		return 0, fmt.Errorf("invalid log level '%s'", s)
	}
}
