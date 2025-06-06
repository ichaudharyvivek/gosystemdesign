package states

import (
	"fmt"
	m "lld/vendingmachine/models"
)

type HasMoneyState struct {
	mx Machine
}

func NewHasMoneyState(machine Machine) *HasMoneyState {
	return &HasMoneyState{mx: machine}
}

func (s *HasMoneyState) InsertCoins(coins []m.Coin) {
	fmt.Println("HasMoneyState: InsertCoins")

	sum := 0
	for _, c := range coins {
		sum += int(c)
	}

	s.mx.AddBalance(sum)
}

func (s *HasMoneyState) SelectProduct(code string) {
	fmt.Println("HasMoneyState: SelectProduct")
	balance := s.mx.GetBalance()

	value, found := s.mx.GetProduct(code)
	if !found {
		fmt.Println("Item not found")
		fmt.Println("Refunding:", balance)
	}

	if value.Price > balance {
		fmt.Println("Insufficient balance")
		fmt.Println("Refunding: ", balance)
	}

	s.mx.SetState(NewDispensingState(s.mx))
}

func (s *HasMoneyState) Dispense() {
	fmt.Println("Please select the product first")
}

func (s *HasMoneyState) OutOfOrder() {
	fmt.Println("Machine is working, please insert the coin and select the product")
}
