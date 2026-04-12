package formatter

import (
	"encoding/json"
	"lld-logger/internal/model"
	"time"
)

type JSONFormatter struct{}

func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

func (f *JSONFormatter) Format(entry model.Record) []byte {
	data := make(map[string]any)

	data["level"] = entry.Level.String()
	data["fields"] = entry.Fields
	data["message"] = entry.Message
	data["timestamp"] = entry.Timestamp.Format(time.RFC3339Nano)

	out, _ := json.Marshal(data)
	return append(out, '\n')
}
