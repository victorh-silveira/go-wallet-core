package postgres

import (
	"context"
	"errors"
	"sync"

	"github.com/victor-silveira/go-wallet-core/src/internal/domain/entity"
)

type WalletRepository struct {
	mu           sync.RWMutex
	accounts     map[string]*entity.Account
	transactions map[string][]*entity.Transaction
}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{
		accounts:     make(map[string]*entity.Account),
		transactions: make(map[string][]*entity.Transaction),
	}
}

func (r *WalletRepository) SaveAccount(ctx context.Context, account *entity.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.accounts[account.ID] = account
	return nil
}

func (r *WalletRepository) GetByID(ctx context.Context, id string) (*entity.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	acc, ok := r.accounts[id]
	if !ok {
		return nil, errors.New("account not found")
	}
	return acc, nil
}

func (r *WalletRepository) GetByUserID(ctx context.Context, userID string) (*entity.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, acc := range r.accounts {
		if acc.UserID == userID {
			return acc, nil
		}
	}
	return nil, errors.New("account not found for user")
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, accountID string, amount float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	acc, ok := r.accounts[accountID]
	if !ok {
		return errors.New("account not found")
	}

	if acc.Balance+amount < 0 {
		return errors.New("insufficient balance")
	}
	acc.Balance += amount
	return nil
}

func (r *WalletRepository) SaveTransaction(ctx context.Context, trx *entity.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.transactions[trx.AccountID] = append(r.transactions[trx.AccountID], trx)
	return nil
}

func (r *WalletRepository) FindAllByAccountID(ctx context.Context, accountID string) ([]*entity.Transaction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	trxs, ok := r.transactions[accountID]
	if !ok {
		return []*entity.Transaction{}, nil
	}
	return trxs, nil
}
