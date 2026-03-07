package main

import d "designpatterns/structural/decorator"

func main() {
	base := &d.BaseNotifier{}

	email := &d.EmailNotification{Wrapped: base, Email: "x@y.com"}
	sms := &d.SmsNotification{Wrapped: base, Phone: "1234567890"}

	email.Send("Hello World!")
	sms.Send("Sent SMS")
}
