/*
The State Pattern allows an object to alter its behavior when its internal state changes.
The object will appear to change its class.
Class Diagram:

	+------------------+
	|     State        |
	+------------------+
	| +InsertCoin()    |
	| +SelectProduct() |
	| +Dispense()      |
	+------------------+
			^
			|
		+---------+---------+---------+---------+---------------------+
		|                   |                   |                     |

	+---------------+  +----------------+  +----------------+   +--------------------+
	|   IdleState   |  | HasMoneyState  |  | DispensingState|   |  OutOfStockState   |
	+---------------+  +----------------+  +----------------+   +--------------------+

	+------------------------+
	|   VendingMachine       |
	+------------------------+
	| -currentState: State   |
	| -stock: int            |
	+------------------------+
	| +SetState(s State)     |
	| +<delegates to state>  |
	+------------------------+
*/
package state

// State defines behavior for each state
type State interface {
	InsertCoin()
	SelectProduct()
	Dispense()
}

// VendingMachine is the context
type VendingMachine struct {
	CurrentState State
	stock        int
}

func NewVendingMachine(stock int) *VendingMachine {
	return &VendingMachine{stock: stock}
}

// SetState allows transition
func (v *VendingMachine) SetState(state State) {
	v.CurrentState = state
}
