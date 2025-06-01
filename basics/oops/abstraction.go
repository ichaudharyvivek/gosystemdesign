// go:build abstraction
package main

import (
	"fmt"
)

// Define the abstraction (interface)
type PaymentProcessor interface {
	Pay(amount float64) error
}

// Concrete implementation - Credit Card
type CreditCardProcessor struct {
	CardNumber string
}

func (cc CreditCardProcessor) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using Credit Card [%s]\n", amount, cc.CardNumber)
	return nil
}

// Concrete implementation - PayPal
type PayPalProcessor struct {
	Email string
}

func (pp PayPalProcessor) Pay(amount float64) error {
	fmt.Printf("Paid %.2f using PayPal account [%s]\n", amount, pp.Email)
	return nil
}

// Use the abstraction
func Checkout(p PaymentProcessor, amount float64) {
	if err := p.Pay(amount); err != nil {
		fmt.Println("Payment failed:", err)
	}
}

func main() {
	// Using Credit Card
	cc := CreditCardProcessor{CardNumber: "1234-5678-9012-3456"}
	Checkout(cc, 100.00)

	// Using PayPal
	pp := PayPalProcessor{Email: "user@example.com"}
	Checkout(pp, 59.99)

	// NOTE: All concrete implementation are dynamically implementing the interface.
}
