package encoder

import "math/rand/v2"

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
