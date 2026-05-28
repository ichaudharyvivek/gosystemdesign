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

func (b *EventBus) UnSubscribe(observer Observer) {
	for i, o := range b.observers {
		if o == observer {
			b.observers = append(b.observers[:i], b.observers[i+1:]...)
			break
		}
	}
}

func (b *EventBus) Notify(event string, data any) {
	for _, o := range b.observers {
		o.Update(event, data)
	}
}
