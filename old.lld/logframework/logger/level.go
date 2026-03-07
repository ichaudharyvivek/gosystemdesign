// log_level
package logger

type LogLevel string

const (
	TRACE   LogLevel = "TRACE"
	DEBUG   LogLevel = "DEBUG"
	INFO    LogLevel = "INFO"
	WARNING LogLevel = "WARNING"
	ERROR   LogLevel = "ERROR"
	FATAL   LogLevel = "FATAL"
)

func (l LogLevel) Ordinal() int {
	switch l {
	case TRACE:
		return 0
	case DEBUG:
		return 1
	case INFO:
		return 2
	case WARNING:
		return 3
	case ERROR:
		return 4
	case FATAL:
		return 5
	default:
		return 999
	}
}
