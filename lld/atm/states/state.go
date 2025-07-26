package states

import "lld/atm/models"

type Machine interface {
	SetState(state State)
	GetState() State

	GetCurrentCard() *models.Card
	SetCurrentCard(card *models.Card) error
	ClearCurrentCard()

	WithdrawCash(amount int) error
	DepositCash(amount int) error
	CheckBalance() (float64, error)
}

type State interface {
	InsertCard(ctx Machine, card *models.Card) error
	InsertPin(ctx Machine, pin string) error
	SelectTransaction(ctx Machine, txnType models.TransactionType, amount ...int) error
	EjectCard(ctx Machine) error
}
