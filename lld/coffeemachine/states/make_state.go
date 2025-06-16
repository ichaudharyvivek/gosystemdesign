package states

import (
	"fmt"
	"sync"
)

type MakeState struct {
	mx Machine
	mu sync.Mutex
}

func NewMakeState(mx Machine) *MakeState {
	return &MakeState{mx: mx}
}

func (s *MakeState) InsertMoney(amount int) {
	fmt.Println("Coffee Machine currently making coffee")
}

func (s *MakeState) Select(id string) {
	fmt.Println("Coffee Machine currently making coffee")
}

func (s *MakeState) Brew() {
	s.mu.Lock()
	defer s.mu.Unlock()

	refund, err := s.mx.BrewCoffee()
	if err != nil {
		fmt.Println("Error encountered. Exiting")
	}

	fmt.Println("Done, coffee brewed")
	if refund > 0 {
		fmt.Println("Balance returned: ", refund)
	}
}
