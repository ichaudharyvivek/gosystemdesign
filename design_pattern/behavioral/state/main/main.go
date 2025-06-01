package main

import (
	s "design-patterns/behavioral/state"
	"fmt"
)

func main() {
	Vm := s.NewVendingMachine(2)
	Vm.SetState(&s.IdleState{Vm: Vm})

	Vm.CurrentState.InsertCoin()
	Vm.CurrentState.SelectProduct()
	Vm.CurrentState.Dispense()

	fmt.Println("---")

	Vm.CurrentState.InsertCoin()
	Vm.CurrentState.SelectProduct()
	Vm.CurrentState.Dispense()

	fmt.Println("---")

	Vm.CurrentState.InsertCoin()
}
