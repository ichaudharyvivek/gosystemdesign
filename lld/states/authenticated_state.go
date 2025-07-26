package states

import (
	"fmt"
	"lld/atm/models"
)

type AuthenticatedState struct{}

func NewAuthenticatedState() *AuthenticatedState {
	return &AuthenticatedState{}
}

func (s *AuthenticatedState) InsertCard(ctx Machine, card *models.Card) error {
	return fmt.Errorf("card already inserted")
}

func (s *AuthenticatedState) InsertPin(ctx Machine, pin string) error {
	return fmt.Errorf("already authenticated")
}

func (s *AuthenticatedState) SelectTransaction(ctx Machine, txnType models.TransactionType, args ...int) error {
	fmt.Println("== processing transaction ==")
	switch txnType {
	case models.TransactionCredit:
		if len(args) != 1 {
			return fmt.Errorf("invalid arguments provided")
		}

		amount := args[0]
		ctx.DepositCash(amount)
	case models.TransactionDebit:
		if len(args) != 1 {
			return fmt.Errorf("invalid arguments provided")
		}

		amount := args[0]
		ctx.WithdrawCash(amount)
	case models.TransactionBalanceInquiry:
		if len(args) > 0 {
			return fmt.Errorf("invalid arguments provided")
		}
		ctx.CheckBalance()
	default:
		return fmt.Errorf("invalid operation type")
	}

	return nil
}

func (s *AuthenticatedState) EjectCard(ctx Machine) error {
	ctx.ClearCurrentCard()
	ctx.SetState(NewIdleState())
	fmt.Println("== transaction complete, card ejected ==")
	return nil
}
