package entity

import (
	"errors"
	"time"
)

type Account struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccount(id, userID string) (*Account, error) {
	if id == "" {
		return nil, errors.New("account id is required")
	}
	if userID == "" {
		return nil, errors.New("user id is required")
	}
	now := time.Now()
	return &Account{
		ID:        id,
		UserID:    userID,
		Balance:   0,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (a *Account) UpdateBalance(deltaCentavos int64) error {
	newBalance := a.Balance + deltaCentavos
	if newBalance < 0 {
		return ErrInsufficientBalance
	}
	a.Balance = newBalance
	a.UpdatedAt = time.Now()
	return nil
}
