package models

import (
	"fmt"
	"time"
)

type TransactionType string

const (
	TransactionDebit          TransactionType = "DEBIT"
	TransactionCredit         TransactionType = "CREDIT"
	TransactionBalanceInquiry TransactionType = "BALANCE_INQUIRY"
)

type Transaction struct {
	ID            string
	Type          TransactionType
	Amount        float64
	AccountNumber string
	Timestamp     time.Time
}

func NewTransaction(txnType TransactionType, amount float64, accountNumber string) *Transaction {
	return &Transaction{
		ID:            generateTransactionID(),
		Type:          txnType,
		Amount:        amount,
		AccountNumber: accountNumber,
		Timestamp:     time.Now(),
	}
}

func generateTransactionID() string {
	return fmt.Sprintf("TXN_%d", time.Now().UnixNano())
}
