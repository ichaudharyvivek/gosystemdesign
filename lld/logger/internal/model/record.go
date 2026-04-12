package model

import "time"

type Record struct {
	Level     Level
	Message   string
	Fields    map[string]any
	Timestamp time.Time
}
