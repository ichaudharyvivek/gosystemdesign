/*
The simple idea behind decorator is to wrap an existing class to modify/extend its behavior.
We can simple do interface implementation also, but to be a decorator, we need a Wrapped class.
Diagram:

           +--------------------+
           |    Component       |
           |--------------------|
           | + operation()      |
           +--------+-----------+
                    ^
                    |
      +-------------+-------------+
      |                           |
+-------------+           +----------------+
|  Concrete   |           |   Decorator    |
| Component   |           |----------------|
|-------------|           | - component    | // Notice how we wrapped the component inside
| +operation()|           | +operation()   |
+-------------+           +-------+--------+
                                  ^
                                  |
              +-------------------+------------------+
              |                                      |
      +--------------------+               +--------------------+
      | ConcreteDecoratorA |               | ConcreteDecoratorB |
      |--------------------|               |--------------------|
      | +operation()       |               | +operation()       |
      +--------------------+               +--------------------+
*/

package decorator

import (
	"fmt"
)

// Interface definition
type Notifier interface {
	Send(message string)
}

// Concrete implementation of an interface
type BaseNotifier struct{}

func (b *BaseNotifier) Send(message string) {
	fmt.Println("[Base]", message)
}

// Decorator 1
// Email Notification
type EmailNotification struct {
	Wrapped Notifier
}

func (e *EmailNotification) Send(message string) {
	e.Wrapped.Send(fmt.Sprintf("[Email] %s", message))
}

// Decorator 2
// SMS Notification
type SmsNotification struct {
	Wrapped Notifier
}

func (e *SmsNotification) Send(message string) {
	e.Wrapped.Send(fmt.Sprintf("[SMS] %s", message))
}
