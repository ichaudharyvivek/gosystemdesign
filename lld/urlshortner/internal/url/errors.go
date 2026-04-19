package url

import "errors"

var (
	ErrNotFound      = errors.New("url not found")
	ErrAlreadyExists = errors.New("url already exists")
	ErrInvalidInput  = errors.New("invalid input")
)
