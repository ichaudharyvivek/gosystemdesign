package states

import (
	"fmt"
	"sync"
)

type IdleState struct {
	mx Machine
	mu sync.Mutex
}

func NewIdleState(mx Machine) *IdleState {
	return &IdleState{mx: mx}
}

func (s *IdleState) InsertMoney(amount int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.mx.AddBalance(amount)
	s.mx.SetState(NewHasMoneyState(s.mx))

	fmt.Println("Added money:", amount)
}

func (s *IdleState) Select(id string) {
	fmt.Println("Please insert coin to continue")
}

func (s *IdleState) Brew() {
	fmt.Println("Please insert coin to continue")
}
