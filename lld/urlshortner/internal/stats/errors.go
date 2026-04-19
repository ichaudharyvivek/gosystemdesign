package stats

import "errors"

var (
	ErrNotFound     = errors.New("stats not found")
	ErrInvalidInput = errors.New("invalid input")
)
