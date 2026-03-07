package services

import (
	"fmt"
	"lld/atm/models"
	"sync"
)

type BankService struct {
	cards            map[string]*models.Card
	accounts         map[string]*models.Account
	cardToAccountMap map[*models.Card]*models.Account
	mu               sync.RWMutex
}

func NewBankService() *BankService {
	return &BankService{
		cards:            make(map[string]*models.Card),
		accounts:         make(map[string]*models.Account),
		cardToAccountMap: make(map[*models.Card]*models.Account),
	}
}

func (bs *BankService) GetCard(cardNumber string) (*models.Card, error) {
	found, ok := bs.cards[cardNumber]
	if !ok {
		return nil, fmt.Errorf("resource not found")
	}

	return found, nil
}

func (bs *BankService) GetBalance(card *models.Card) (float64, error) {
	acc, ok := bs.cardToAccountMap[card]
	if !ok {
		return 0.0, fmt.Errorf("resource not found")
	}

	return acc.GetBalance(), nil
}

func (bs *BankService) Authenticate(card *models.Card, pin string) error {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	cardNumber := card.GetCardNumber()
	found, ok := bs.cards[cardNumber]
	if !ok {
		return fmt.Errorf("resource not found")
	}

	if len(pin) != 4 {
		return fmt.Errorf("invalid data")
	}

	if !found.ValidatePin(pin) {
		return fmt.Errorf("invalid pin")
	}

	return nil
}

func (bs *BankService) Deposit(card *models.Card, amount float64) error {
	c, ok := bs.cardToAccountMap[card]
	if !ok {
		return fmt.Errorf("resource not found")
	}

	return c.Credit(amount)
}

func (bs *BankService) Withdraw(card *models.Card, amount float64) error {
	c, ok := bs.cardToAccountMap[card]
	if !ok {
		return fmt.Errorf("resource not found")
	}

	return c.Debit(amount)
}

func (bs *BankService) CreateAccount(accountNumber string, initialBalance float64) (*models.Account, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	if len(accountNumber) < 16 {
		return nil, fmt.Errorf("invalid data")
	}

	newAccount := models.NewAccount(accountNumber, initialBalance)
	bs.accounts[accountNumber] = newAccount
	return newAccount, nil
}

func (bs *BankService) CreateCard(cardNumber string, pin string) (*models.Card, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	if len(cardNumber) < 16 {
		return nil, fmt.Errorf("invalid data")
	}

	newCard := models.NewCard(cardNumber, pin)
	bs.cards[cardNumber] = newCard
	return newCard, nil
}

func (bs *BankService) LinkCardToAccount(account *models.Account, card *models.Card) error {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	if account == nil || card == nil {
		return fmt.Errorf("invalid data")
	}

	bs.cardToAccountMap[card] = account
	return nil
}
