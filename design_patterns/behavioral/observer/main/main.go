package main

import o "designpatterns/behavioral/observer"

func main() {
	sms := &o.SmsClient{}
	email := &o.EmailClient{}

	ns := &o.NotificationService{}
	ns.Subscribe(email)
	ns.Subscribe(sms)
	ns.NotifyAll("Hello World!")
}
