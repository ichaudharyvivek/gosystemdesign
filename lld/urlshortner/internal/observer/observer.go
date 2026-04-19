package observer

type Observer interface {
	Update(event string, data any)
}
