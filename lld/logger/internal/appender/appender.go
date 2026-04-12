package appender

type Appender interface {
	Append(data []byte) error
}
