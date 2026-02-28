package observer

type NotificationService struct {
	observers []Observer
}

func (n *NotificationService) Subscribe(o Observer) {
	n.observers = append(n.observers, o)
}

func (n *NotificationService) Unsubscribe(o Observer) {
	for i, obs := range n.observers {
		if obs == o {
			// OR USE: n.observers = slices.Delete(n.observers, i, i+1)
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NotificationService) NotifyAll(message string) {
	for _, obs := range n.observers {
		obs.Update(message)
	}
}
