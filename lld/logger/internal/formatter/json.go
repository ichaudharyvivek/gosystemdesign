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

func (f *JSONFormatter) Format(record model.Record) []byte {
	data := make(map[string]any)

	data["level"] = record.Level.String()
	data["fields"] = record.Fields
	data["message"] = record.Message
	data["timestamp"] = record.Timestamp.Format(time.RFC3339Nano)

	out, _ := json.Marshal(data)
	return append(out, '\n')
}
