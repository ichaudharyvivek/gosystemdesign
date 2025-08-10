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

type Dispensor interface {
	SetNext(next Dispensor)
	Dispense(amount int)
}

type BaseDispensor struct {
	denomination int
	count        int
	next         Dispensor
}

func NewBaseDispensor(denomination, count int) *BaseDispensor {
	return &BaseDispensor{
		denomination: denomination,
		count:        count,
	}
}

func (d *BaseDispensor) SetNext(next Dispensor) {
	d.next = next
}

func (d *BaseDispensor) Dispense(amt int) {
	
}
