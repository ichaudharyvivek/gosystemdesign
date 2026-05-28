package encoder

type Base62Encoder struct{}

func NewBase62Encoder() *Base62Encoder {
	return &Base62Encoder{}
}

func (e *Base62Encoder) Encode(id int64) string {
	char := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if id == 0 {
		return "0"
	}

	var encoded []byte
	for id > 0 {
		remainder := id % 62
		encoded = append([]byte{char[remainder]}, encoded...)
		id = id / 62
	}

	return string(encoded)
}
