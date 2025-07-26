package models

import (
	"fmt"
	"lld/atm/dispenser"
)

type CashDispenser struct {
	dispenserChain dispenser.Dispenser
}

func NewCashDispenser() *CashDispenser {
	// Create dispenser chain
	dispenser2000 := dispenser.NewNoteDispenser(2000, 10)
	dispenser500 := dispenser.NewNoteDispenser(500, 10)
	dispenser100 := dispenser.NewNoteDispenser(100, 20)

	// Set up chain
	dispenser2000.SetNext(dispenser500)
	dispenser500.SetNext(dispenser100)

	return &CashDispenser{
		dispenserChain: dispenser2000,
	}
}

func (cd *CashDispenser) DispenseCash(amount int) error {
	if !cd.CanDispense(amount) {
		return fmt.Errorf("ATM can only dispense multiples of 100")
	}

	return cd.dispenserChain.Dispense(amount)
}

func (cd *CashDispenser) CanDispense(amount int) bool {
	return amount > 0 && amount%100 == 0
}
