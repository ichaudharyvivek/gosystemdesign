package services

import (
	"fmt"
	"lld/atm/models"
	"lld/atm/states"
	"sync"
)

var (
	instance *ATM
	once     sync.Once
)

type ATM struct {
	currentState  states.State
	currentCard   *models.Card
	bankService   *BankService
	cashDispenser *models.CashDispenser
	mu            sync.RWMutex
}

func GetATMInstance(bs *BankService) *ATM {
	once.Do(func() {
		instance = &ATM{
			bankService:   bs,
			cashDispenser: models.NewCashDispenser(),
			currentState:  states.NewIdleState(),
		}
	})

	return instance
}

// API methods
func (mx *ATM) InsertCard(card *models.Card) (_ error) {
	return mx.currentState.InsertCard(mx, card)
}

func (mx *ATM) InsertPin(pin string) (_ error) {
	return mx.currentState.InsertPin(mx, pin)
}

func (mx *ATM) SelectTransaction(txnType models.TransactionType, amount ...int) (_ error) {
	return mx.currentState.SelectTransaction(mx, txnType, amount...)
}

func (mx *ATM) EjectCard() (_ error) {
	return mx.currentState.EjectCard(mx)
}

// Internal methods
func (mx *ATM) SetState(state states.State) {
	mx.currentState = state
}

func (mx *ATM) GetState() states.State {
	return mx.currentState
}

func (mx *ATM) GetCurrentCard() *models.Card {
	return mx.currentCard
}

func (mx *ATM) SetCurrentCard(card *models.Card) error {
	if card == nil {
		return fmt.Errorf("invalid data")
	}

	mx.currentCard = card
	return nil
}

func (mx *ATM) ClearCurrentCard() {
	mx.currentCard = nil
}

func (mx *ATM) WithdrawCash(amount int) error {
	if mx.currentCard == nil {
		return fmt.Errorf("insert card")
	}

	mx.cashDispenser.DispenseCash(amount)
	return mx.bankService.Withdraw(mx.currentCard, float64(amount))
}

func (mx *ATM) DepositCash(amount int) error {
	if mx.currentCard == nil {
		return fmt.Errorf("insert card")
	}

	return mx.bankService.Deposit(mx.currentCard, float64(amount))
}

func (mx *ATM) CheckBalance() (float64, error) {
	if mx.currentCard == nil {
		return 0.0, fmt.Errorf("insert card")
	}

	if _, err := mx.bankService.GetBalance(mx.currentCard); err != nil {
		return 0.0, err
	}

	fmt.Printf("Balance: %d\n", 100)
	return 0.0, nil
}
