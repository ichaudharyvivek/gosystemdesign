package states

import (
	"fmt"
	m "lld/vendingmachine/models"
)

type HasMoneyState struct {
	mx Machine
}

func NewHasMoneyState(mx Machine) *HasMoneyState {
	return &HasMoneyState{mx}
}

func (s *HasMoneyState) InsertCoins(coins []m.Coin) {
	fmt.Println("Has money state, inserting more coins")

	total := 0
	for _, c := range coins {
		total += int(c)
	}
	s.mx.AddBalance(total)
}

func (s *HasMoneyState) SelectProduct(code string) {
	fmt.Println("Initiated product selection")

	paid := s.mx.GetBalance()
	pd, found := s.mx.GetProductByCode(code)
	if !found {
		fmt.Println("Product not found, please select again.")
		return
	}

	if !s.mx.IsProductOutOfStock(code) {
		fmt.Println("Product out of stock. Refunding amount")
		s.mx.Refund(s.mx.GetBalance())
		s.mx.SetState(NewIdleState(s.mx))
	}

	if pd.Price > paid {
		fmt.Println("Insufficient balance, please add more coins")
		return
	}

	s.mx.SetSelectedProduct(code)
	s.mx.SetState(NewDispensingState(s.mx))
	fmt.Println("Product selected.", pd)
}

func (s *HasMoneyState) Dispense() {
	fmt.Println("Please select product to continue")
}
