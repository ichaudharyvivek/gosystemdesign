package states

import (
	"fmt"
	m "lld/vendingmachine/models"
)

type DispensingState struct {
	mx Machine
}

func NewDispensingState(mx Machine) *DispensingState {
	return &DispensingState{mx}
}

func (s *DispensingState) InsertCoins(coins []m.Coin) {
	fmt.Println("Processing product and dispensing, please wait...")
}

func (s *DispensingState) SelectProduct(code string) {
	fmt.Println("Processing product and dispensing, please wait...")
}

func (s *DispensingState) Dispense() {
	fmt.Println("Dispensing selected product")

	refund, ok := s.mx.DispenseProduct()
	if !ok {
		fmt.Println("Something went wrong. Initiating refund...")
		s.mx.Refund(s.mx.GetBalance())
	}

	if refund > 0 {
		s.mx.Refund(refund)
	}
	s.mx.SetState(NewIdleState(s.mx))
}
