// go:build dip

/*
Why this follows Dependency Injection Principle (DIP):
  - Notification depends on the MessageService interface, not a concrete EmailService.
  - You can swap in SMSService, SlackService, etc., without changing Notification.

# High-level logic is decoupled from low-level implementations.
# Equivalent java code:
```java

	interface MessageService {
	    void sendMessage(String msg);
	}

	class EmailService implements MessageService {
	    public void sendMessage(String msg) {
	        System.out.println("Sending email: " + msg);
	    }
	}

	class Notification {
	    private MessageService service;

	    public Notification(MessageService service) {
	        this.service = service;
	    }

	    public void notifyUser() {
	        service.sendMessage("Hello User!");
	    }
	}

```
# Diagram:
```

	 --- Main Class [Notification]--------
	 --- > MessageService service --------
	              |
				  |
				  *
	 --- Interface MessageService --------
	 --- > sendMessage(String message) ---
				  *
				  |
				  |
	 --- EmailService --------------------
	 --- impl. MessageService ------------
	 --- > notify() ----------------------

```
*/
package main

import "fmt"

// Abstraction
type MessageService interface {
	sendMessage(message string)
}

// Low level - Impl details depends on abstraction (interface)
type EmailService struct{}

func (e *EmailService) sendMessage(message string) {
	fmt.Println("Sending Email:", message)
}

// Low level - Impl details depends on abstraction (interface)
type SMSService struct{}

func (s *SMSService) sendMessage(message string) {
	fmt.Println("Sending SMS: ", message)
}

// High level - depends on abstraction (interface)
type Notification struct {
	service MessageService
}

func (n *Notification) NotifyUser() {
	n.service.sendMessage("Hello User")
}
