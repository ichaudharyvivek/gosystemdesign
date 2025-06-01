/*
The Strategy Pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable.
The strategy lets the algorithm vary independently from clients that use it.
EG: Imagine a payment system that supports multiple payment methods: Credit Card, PayPal, UPI.

	The strategy is the selected payment algorithm.

Diagram:

	+--------------------+
	|  PaymentStrategy   |<------------------------+
	+--------------------+                         |
	| +Pay(amount)       |                         |
	+--------------------+                         |
	       ^                                       ^
	       |                                       |

	+------------------+             +---------------------+
	| CreditCardPayment|             |     PayPalPayment   |
	+------------------+             +---------------------+
	| +Pay(amount)     |             | +Pay(amount)        |
	+------------------+             +---------------------+

	       used by
	           |
	           v
	+----------------------+
	|   PaymentProcessor   |
	+----------------------+
	| - strategy           |
	| +SetStrategy()       |
	| +PayAmount()         |
	+----------------------+
*/
package strategy

import "fmt"

// PaymentStrategy refers the methods all various sets of algorithm is gonna use
type PaymentStrategy interface {
	Pay(amount int)
}

// Strategy 1
type CreditCardPayment struct {
}

func (p *CreditCardPayment) Pay(amount int) {
	fmt.Println("Paying via credit card. Amount:", amount)
}

// Strategy 2
type DebitCardPayment struct {
}

func (p *DebitCardPayment) Pay(amount int) {
	fmt.Println("Paying via debit card. Amount:", amount)
}
