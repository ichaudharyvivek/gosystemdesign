/*
Chain of Responsibility is a behavioral design pattern where you pass a request along a chain of handlers,
Each handler decides:
 1. To process it, or
 2. To pass it along to the next handler in the chain.

This decouple the sender of a request from the receiver.
The sender doesn’t need to know which object will handle the request.
Diagram:

	           +-------------------+
	           |   <<interface>>   |
	           |   DispenseChain   |
	           +-------------------+
	           | + SetNext(c)      |
	           | + Dispense(amt)   |
	           +-------------------+
	                    ^
	                    |
	           +-----------------------+
	           |  BaseDispenser        |
	           +-----------------------+
	           | - next: DispenseChain |
	           +-----------------------+
	           | + SetNext(c)          |
	           | + CallNext(amt)       |
	           +-----------------------+
	                    ^
						|
	+-------------------+------+-------------+-----+
	|                          |                   |

+----------------+    +----------------+   +----------------+
| Rs2000Dispenser|    | Rs500Dispenser |   | Rs100Dispenser |
+----------------+    +----------------+   +----------------+
| (BaseDispenser)|    | (BaseDispenser)|   | (BaseDispenser)|
+----------------+    +----------------+   +----------------+
| + Dispense()   |    | + Dispense()   |   | + Dispense()   |
+----------------+    +----------------+   +----------------+
*/
package chain

import "fmt"

type Dispensor interface {
	SetNext(next Dispensor)
	Dispense(amount int)
}

type BaseDispensor struct {
	denomination int
	count        int
	next         Dispensor
}

// Creates a new dispenser for a specific denomination with available note count
func NewBaseDispensor(denomination, count int) *BaseDispensor {
	return &BaseDispensor{
		denomination: denomination,
		count:        count,
	}
}

// Links the current dispenser to the next one in the chain
func (d *BaseDispensor) SetNext(next Dispensor) {
	d.next = next
}

// Dispenses notes of current denomination, passes remainder to the next dispenser
func (d *BaseDispensor) Dispense(amt int) {
	if amt <= 0 || amt%100 != 0 {
		fmt.Printf("Cannot dispense amount: %d\n", amt)
		return
	}

	if amt >= d.denomination {
		numOfNotes := amt / d.denomination
		dispenseAmount := numOfNotes * d.denomination
		remaining := amt - dispenseAmount

		if numOfNotes > d.count {
			fmt.Println("Not enough notes")
			return
		}

		if remaining > 0 {
			if d.next != nil {
				d.next.Dispense(remaining)
			} else {
				fmt.Println("Cannot print amount")
				return
			}
		}

		d.count -= numOfNotes
		fmt.Printf("Dispensing %dx%d = %d\n", numOfNotes, d.denomination, dispenseAmount)

	} else {
		d.next.Dispense(amt)
	}
}
