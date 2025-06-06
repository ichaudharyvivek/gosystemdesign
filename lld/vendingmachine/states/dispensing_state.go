package states

import (
	"fmt"
	m "lld/vendingmachine/models"
)

type DispensingState struct {
	mx Machine
}

func NewDispensingState(machine Machine) *DispensingState {
	return &DispensingState{mx: machine}
}

func (s *DispensingState) InsertCoins(coins []m.Coin) {
	fmt.Println("Product selected, dispensing product")
}

func (s *DispensingState) SelectProduct(code string) {
	fmt.Println("Product selected, dispensing product")
}

func (s *DispensingState) Dispense() {
	fmt.Println("DispensingState: Dispense")
	s.mx.SetState(NewIdleState(s.mx))
}

func (s *DispensingState) OutOfOrder() {
	fmt.Println("Machine is working, dispensing product")
}
