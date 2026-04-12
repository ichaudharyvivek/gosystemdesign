package appender

import "lld-logger/internal/model"

type Appender interface {
	Append(record model.Record) error
	Close() error
}
