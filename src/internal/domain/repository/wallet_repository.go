package repository

import (
	"context"

	"github.com/victor-silveira/go-wallet-core/src/internal/domain/entity"
)

type AccountRepository interface {
	SaveAccount(ctx context.Context, account *entity.Account) error
	GetByID(ctx context.Context, id string) (*entity.Account, error)
	GetByUserID(ctx context.Context, userID string) (*entity.Account, error)
	UpdateBalance(ctx context.Context, accountID string, amount float64) error
}

type TransactionRepository interface {
	SaveTransaction(ctx context.Context, transaction *entity.Transaction) error
	FindAllByAccountID(ctx context.Context, accountID string) ([]*entity.Transaction, error)
}
