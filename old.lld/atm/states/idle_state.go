package states

import (
	"fmt"
	"lld/atm/models"
)

type IdleState struct{}

func NewIdleState() *IdleState {
	return &IdleState{}
}

func (s *IdleState) InsertCard(ctx Machine, card *models.Card) error {
	if card == nil {
		return fmt.Errorf("invalid card")
	}

	ctx.SetCurrentCard(card)
	ctx.SetState(NewCardInsertedState())
	fmt.Println("== card inserted ==")
	return nil
}

func (s *IdleState) InsertPin(ctx Machine, pin string) error {
	return fmt.Errorf("please insert card")
}

func (s *IdleState) SelectTransaction(ctx Machine, txnType models.TransactionType, amount ...int) error {
	return fmt.Errorf("please insert card")
}

func (s *IdleState) EjectCard(ctx Machine) error {
	return fmt.Errorf("please insert card")
}
