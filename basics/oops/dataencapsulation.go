// go:build encapsulation
package main

import (
	"errors"
	"fmt"
)

// Encapsulated struct
type BankAccount struct {
	accountHolder string  // private (unexported)
	balance       float64 // private (unexported)
}

// Constructor
func NewBankAccount(holder string, initialDeposit float64) (*BankAccount, error) {
	if initialDeposit < 0 {
		return nil, errors.New("initial deposit cannot be negative")
	}
	return &BankAccount{
		accountHolder: holder,
		balance:       initialDeposit,
	}, nil
}

// Getter (public)
func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

// Deposit method (public)
func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	b.balance += amount
	return nil
}

// Withdraw method (public)
func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be positive")
	}
	if amount > b.balance {
		return errors.New("insufficient funds")
	}
	b.balance -= amount
	return nil
}

func main() {
	account, err := NewBankAccount("Alice", 1000)
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}

	fmt.Println("Initial balance:", account.GetBalance())

	_ = account.Deposit(500)
	fmt.Println("After deposit:", account.GetBalance())

	err = account.Withdraw(300)
	if err != nil {
		fmt.Println("Withdraw error:", err)
	} else {
		fmt.Println("After withdrawal:", account.GetBalance())
	}
}
