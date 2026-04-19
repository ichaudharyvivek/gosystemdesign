package encoder

const (
	MAX_LENGTH = 7
	CHARACTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type Encoder interface {
	Generate() (string, error)
}
