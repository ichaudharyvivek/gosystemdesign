package logger

import (
	"lld-logger/internal/appender"
	"lld-logger/internal/formatter"
	"lld-logger/internal/model"
	"os"
)

type Logger struct {
	level     model.Level
	formatter formatter.Formatter
	appenders []appender.Appender
}

func New() *Logger {
	level := model.InfoLevel
	formatter := formatter.NewTextFormatter()
	appenders := []appender.Appender{appender.NewConsoleAppender(os.Stdout)}

	return &Logger{
		level:     level,
		formatter: formatter,
		appenders: appenders,
	}
}

func (l *Logger) SetLevel(level model.Level) {
	l.level = level
}

func (l *Logger) SetAppenders(appenders []appender.Appender) {
	l.appenders = appenders
}

func (l *Logger) SetFormatter(formatter formatter.Formatter) {
	l.formatter = formatter
}

func (l *Logger) Debug() *Entry {
	return l.log(model.DebugLevel)
}

func (l *Logger) Info() *Entry {
	return l.log(model.InfoLevel)
}

func (l *Logger) Warn() *Entry {
	return l.log(model.WarnLevel)
}

func (l *Logger) Error() *Entry {
	return l.log(model.ErrorLevel)
}

func (l *Logger) Fatal() *Entry {
	return l.log(model.FatalLevel)
}

func (l *Logger) log(level model.Level) *Entry {
	return &Entry{
		logger: l,
		level:  level,
	}
}
