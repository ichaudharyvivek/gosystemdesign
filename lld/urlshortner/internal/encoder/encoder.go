package encoder

type Encoder interface {
	Encode(id int64) string
}
