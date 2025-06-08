// log_message
package logger

import (
	"fmt"
	"time"
)

type LogMessage struct {
	Time    time.Time
	Message string
	Level   LogLevel
}

func (l LogMessage) String() string {
	return fmt.Sprintf("[%s] [%s] %s", l.Time.Format(time.RFC3339), l.Level, l.Message)
}
