package logger

import (
	"sync"
	"time"
)

var (
	instance *Logger
	once     sync.Once
)

type Logger struct {
	config *LogConfig
	mu     sync.Mutex
}

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{}
	})

	return instance
}

func (l *Logger) SetConfiguration(cfg *LogConfig) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config = cfg

	// Start all appenders to ensure they're ready to receive logs (e.g. open files, DB connections, etc.)
	// for _, app := range cfg.Appenders {
	// 	app.Start()
	// }
}

func (l *Logger) log(level LogLevel, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if level.Ordinal() < l.config.Level.Ordinal() {
		return
	}

	ln := &LogMessage{
		Message: msg,
		Level:   level,
		Time:    time.Now(),
	}

	for _, a := range l.config.Appenders {
		a.Append(ln)
	}
}

func (l *Logger) Trace(msg string)   { l.log(TRACE, msg) }
func (l *Logger) Debug(msg string)   { l.log(DEBUG, msg) }
func (l *Logger) Info(msg string)    { l.log(INFO, msg) }
func (l *Logger) Warning(msg string) { l.log(WARNING, msg) }
func (l *Logger) Error(msg string)   { l.log(ERROR, msg) }
func (l *Logger) Fatal(msg string)   { l.log(FATAL, msg) }
