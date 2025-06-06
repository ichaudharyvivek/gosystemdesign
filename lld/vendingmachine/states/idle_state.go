package states

import (
	"fmt"
	m "lld/vendingmachine/models"
)

type IdleState struct {
	mx Machine
}

func NewIdleState(mx Machine) *IdleState {
	return &IdleState{mx}
}

func (s *IdleState) InsertCoins(coins []m.Coin) {
	fmt.Println("Idle state, inserting coins")

	amount := 0
	for _, c := range coins {
		amount += int(c)
	}

	s.mx.AddBalance(amount)
	s.mx.SetState(NewHasMoneyState(s.mx))
}

func (s *IdleState) SelectProduct(code string) {
	fmt.Println("Cannot select product. \nPlease insert a coin to continue")
}

func (s *IdleState) Dispense() {
	fmt.Println("Cannot dispense product. \nPlease insert a coin to continue")
}
