package formatter

import "lld-logger/internal/model"

type Formatter interface {
	Format(entry model.Record) []byte
}
