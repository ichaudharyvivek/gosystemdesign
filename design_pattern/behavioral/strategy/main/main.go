package main

import s "design-patterns/behavioral/strategy"

func main() {
	credit := s.CreditCardPayment{}
	paymentProcessor := s.NewPaymentProcessor(&credit)

	paymentProcessor.PayAmount(5000)
}
