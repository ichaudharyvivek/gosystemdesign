package url

import "math/rand/v2"

const (
	MAX_LENGTH = 7
	CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type Encoder interface {
	Generate() (string, error)
}

type BasicEncoder struct{}

func NewBasicEncoder() *BasicEncoder {
	return &BasicEncoder{}
}

func (e *BasicEncoder) Generate() (string, error) {
	b := make([]byte, MAX_LENGTH)
	for i := range b {
		b[i] = CHARACTERS[rand.N(len(CHARACTERS))]
	}

	return string(b), nil
}
