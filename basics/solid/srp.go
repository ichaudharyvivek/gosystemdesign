// go:build srp

/*
Why this follows Single Responsibility Principle (SRP):
  - Invoice handles data storage (just the amount).
  - InvoicePrinter handles presentation (how it’s printed).
  - Each has a single reason to change:

# Change Invoice if invoice data changes
# Change InvoicePrinter if printing format changes
# Clean separation of concerns.
# Equivalent java code:
```java

	class Invoice {
	    private double amount;
	    public Invoice(double amount) { this.amount = amount; }
	    public double getAmount() { return amount; }
	}

	class InvoicePrinter {
	    public void print(Invoice invoice) {
	        System.out.println("Invoice amount: $" + invoice.getAmount());
	    }
	}

```
*/
package main

import "fmt"

type Invoice struct {
	Amount float64
}

type InvoicePrinter struct{}

func (p InvoicePrinter) Print(invoice Invoice) {
	fmt.Printf("Invoice amount: $%.2f\n", invoice.Amount)
}
