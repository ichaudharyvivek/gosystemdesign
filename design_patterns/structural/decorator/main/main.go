package main

import d "design-patterns/structural/decorator"

func main() {
	base := &d.BaseNotifier{}

	email := &d.EmailNotification{Wrapped: base}
	sms := &d.SmsNotification{Wrapped: base}

	email.Send("Hello World!")
	sms.Send("Sent SMS")
}
