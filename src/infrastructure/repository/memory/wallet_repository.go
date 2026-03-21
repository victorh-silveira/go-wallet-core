package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
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
	if account == nil {
		return errors.New("account is required")
	}
	accountCopy := *account
	r.accounts[account.ID] = &accountCopy
	return nil
}

func (r *WalletRepository) GetByID(ctx context.Context, id string) (*entity.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	acc, ok := r.accounts[id]
	if !ok {
		return nil, errors.New("account not found")
	}
	accountCopy := *acc
	return &accountCopy, nil
}

func (r *WalletRepository) GetByUserID(ctx context.Context, userID string) (*entity.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, acc := range r.accounts {
		if acc.UserID == userID {
			accountCopy := *acc
			return &accountCopy, nil
		}
	}
	return nil, errors.New("account not found for user")
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, accountID string, deltaCentavos int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	acc, ok := r.accounts[accountID]
	if !ok {
		return errors.New("account not found")
	}

	if acc.Balance+deltaCentavos < 0 {
		return entity.ErrInsufficientBalance
	}
	acc.Balance += deltaCentavos
	return nil
}

func (r *WalletRepository) SaveTransaction(ctx context.Context, trx *entity.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if trx == nil {
		return errors.New("transaction is required")
	}
	trxCopy := *trx
	r.transactions[trx.AccountID] = append(r.transactions[trx.AccountID], &trxCopy)
	return nil
}

func (r *WalletRepository) FindAllByAccountID(ctx context.Context, accountID string) ([]*entity.Transaction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	trxs, ok := r.transactions[accountID]
	if !ok {
		return []*entity.Transaction{}, nil
	}
	result := make([]*entity.Transaction, 0, len(trxs))
	for _, trx := range trxs {
		trxCopy := *trx
		result = append(result, &trxCopy)
	}
	return result, nil
}
