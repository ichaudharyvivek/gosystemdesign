package states

import (
	"fmt"
	m "lld/vendingmachine/models"
)

type IdleState struct {
	mx Machine // Interface to decouple concrete object
}

func NewIdleState(machine Machine) *IdleState {
	return &IdleState{mx: machine}
}

func (s *IdleState) InsertCoins(coins []m.Coin) {
	fmt.Println("IdleState: InsertCoins")
	sum := 0
	for _, c := range coins {
		sum += int(c)
	}

	s.mx.AddBalance(sum)
	s.mx.SetState(NewHasMoneyState(s.mx))
}

func (s *IdleState) SelectProduct(code string) {
	fmt.Println("Insert coin first before selecting product")
}

func (s *IdleState) Dispense() {
	fmt.Println("Insert coin first before dispensing")
}

func (s *IdleState) OutOfOrder() {
	fmt.Println("Machine is working, insert coin to continue")
}
