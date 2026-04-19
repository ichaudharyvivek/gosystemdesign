package observer

type EventBus struct {
	observers []Observer
}

func NewEventBus() *EventBus {
	return &EventBus{
		observers: make([]Observer, 0),
	}
}

func (b *EventBus) Subscribe(observer Observer) {
	b.observers = append(b.observers, observer)
}

func (b *EventBus) Unsubscribe(observer Observer) {
	for i, o := range b.observers {
		if o == observer {
			b.observers = append(b.observers[:i], b.observers[i+1:]...)
			break
		}
	}
}

func (b *EventBus) Notify(event string, data any) {
	for _, observer := range b.observers {
		observer.Update(event, data)
	}
}
