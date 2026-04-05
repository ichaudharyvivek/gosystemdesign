package formatter

import (
	"encoding/json"
	"lld-logger/internal/core"
	"time"
)

type JSONFormatter struct{}

func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

func (jf *JSONFormatter) Format(entry core.LogEntry) []byte {
	data := make(map[string]any)

	data["timestamp"] = entry.Timestamp.Format(time.RFC3339Nano)
	data["level"] = entry.Level.String()
	data["message"] = entry.Message

	out, _ := json.Marshal(data)
	return append(out, '\n')
}
