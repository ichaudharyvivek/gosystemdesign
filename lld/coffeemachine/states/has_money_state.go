package states

import (
	"fmt"
	"sync"
)

type HasMoneyState struct {
	mx Machine
	mu sync.Mutex
}

func NewHasMoneyState(mx Machine) *HasMoneyState {
	return &HasMoneyState{mx: mx}
}

func (s *HasMoneyState) InsertMoney(amount int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.mx.AddBalance(amount)
	fmt.Println("Added money:", amount)
}

func (s *HasMoneyState) Select(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.mx.SelectCoffee(id)
	if err != nil {
		fmt.Printf("Coffee with '%s' does not exists", id)
		return
	}

	fmt.Println("Product selected:", id)
	s.mx.SetState(NewMakeState(s.mx))
}

func (s *HasMoneyState) Brew() {
	fmt.Println("Please select the product before making coffee")
}
