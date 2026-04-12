package formatter

import "lld-logger/internal/model"

type Formatter interface {
	Format(record model.Record) []byte
}
