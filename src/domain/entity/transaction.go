package entity

import (
	"errors"
	"time"
)

type TransactionType string

const (
	Debit  TransactionType = "DEBIT"
	Credit TransactionType = "CREDIT"
)

type Transaction struct {
	ID          string          `json:"id"`
	AccountID   string          `json:"account_id"`
	Type        TransactionType `json:"type"`
	Amount      float64         `json:"amount"`
	Description string          `json:"description"`
	CreatedAt   time.Time       `json:"created_at"`
}

func NewTransaction(id, accountID, description string, tType TransactionType, amount float64) (*Transaction, error) {
	if id == "" || accountID == "" {
		return nil, errors.New("transaction and account ids are required")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	return &Transaction{
		ID:          id,
		AccountID:   accountID,
		Type:        tType,
		Amount:      amount,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}
