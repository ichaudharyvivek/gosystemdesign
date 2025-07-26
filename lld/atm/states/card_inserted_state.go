package states

import (
	"fmt"
	"lld/atm/models"
)

type CardInsertedState struct{}

func NewCardInsertedState() *CardInsertedState {
	return &CardInsertedState{}
}

func (s *CardInsertedState) InsertCard(ctx Machine, card *models.Card) error {
	return fmt.Errorf("card already inserted")
}

func (s *CardInsertedState) InsertPin(ctx Machine, pin string) error {
	card := ctx.GetCurrentCard()
	if card == nil {
		return fmt.Errorf("invalid data")
	}

	card.ValidatePin(pin)
	ctx.SetState(NewAuthenticatedState())
	fmt.Println("== authentication success ==")
	return nil
}

func (s *CardInsertedState) SelectTransaction(ctx Machine, txnType models.TransactionType, amount ...int) error {
	return fmt.Errorf("please enter your PIN first")
}

func (s *CardInsertedState) EjectCard(ctx Machine) error {
	ctx.ClearCurrentCard()
	ctx.SetState(NewIdleState())
	fmt.Println("Card ejected.")
	return nil
}
