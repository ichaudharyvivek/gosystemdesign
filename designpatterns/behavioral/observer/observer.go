/*
The Observer Pattern defines a one-to-many dependency between objects so that when one object changes state,
all its dependents (observers) are notified and updated automatically.
Diagram:

	    +---------------+
	    |   Observer    |<--------------+
	    +---------------+               |
	    | +Update(msg)  |               |
	    +---------------+               |
	          ^							|
			  |                         |
	+------------------+       +-----------------+
	|   EmailClient    |       |   SMSClient     |
	+------------------+       +-----------------+
	| +Update(msg)     |       | +Update(msg)    |
	+------------------+       +-----------------+

	          notified by
	                |
	                v
	  +----------------------------+
	  |     NotificationService    |
	  +----------------------------+
	  | -observers []Observer      |
	  | +Subscribe(o Observer)     |
	  | +Unsubscribe(o Observer)   |
	  | +NotifyAll(msg string)     |
	  +----------------------------+
*/
package observer

import "fmt"

// Observer interface
type Observer interface {
	Update(message string)
}

// Concrete observer 1 impl Observer
type EmailClient struct{}

func (e *EmailClient) Update(message string) {
	fmt.Println("Email Client updates: ", message)
}

// Concrete observer 2 impl Observer
type SmsClient struct{}

func (e *SmsClient) Update(message string) {
	fmt.Println("SMS Client updates: ", message)
}
