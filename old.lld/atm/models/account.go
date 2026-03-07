package models

import (
	"fmt"
	"sync"
)

type Account struct {
	accountNumber string
	balance       float64
	transactions  []*Transaction
	mu            sync.RWMutex
}

func NewAccount(accountNumber string, balance float64) *Account {
	return &Account{
		accountNumber: accountNumber,
		balance:       balance,
	}
}

func (a *Account) GetAccountNumber() string {
	return a.accountNumber
}

func (a *Account) GetBalance() float64 {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.balance
}

func (a *Account) Debit(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if amount <= 0 {
		return fmt.Errorf("invalid amount")
	}
	if amount > a.balance {
		return fmt.Errorf("insufficient amount")
	}

	a.balance -= amount
	a.transactions = append(a.transactions, NewTransaction(TransactionDebit, amount, a.accountNumber))
	return nil
}

func (a *Account) Credit(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if amount <= 0 {
		return fmt.Errorf("invalid amount")
	}

	a.balance += amount
	a.transactions = append(a.transactions, NewTransaction(TransactionCredit, amount, a.accountNumber))
	return nil
}

func (a *Account) String() string {
	return fmt.Sprintf("Account{accountNumber: %s, balance: %.2f}",
		a.accountNumber, a.balance)
}
