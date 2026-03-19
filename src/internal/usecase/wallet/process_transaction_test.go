package wallet

import (
	"context"
	"errors"
	"testing"

	"github.com/victor-silveira/go-wallet-core/src/internal/domain/entity"
)

type accountRepoStub struct {
	account *entity.Account
	err     error
}

func (s *accountRepoStub) SaveAccount(ctx context.Context, account *entity.Account) error {
	s.account = account
	return nil
}

func (s *accountRepoStub) GetByID(ctx context.Context, id string) (*entity.Account, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.account, nil
}

func (s *accountRepoStub) GetByUserID(ctx context.Context, userID string) (*entity.Account, error) {
	return nil, errors.New("not implemented")
}

func (s *accountRepoStub) UpdateBalance(ctx context.Context, accountID string, amount float64) error {
	return errors.New("not implemented")
}

type txRepoStub struct {
	err error
}

func (s *txRepoStub) SaveTransaction(ctx context.Context, transaction *entity.Transaction) error {
	return s.err
}

func (s *txRepoStub) FindAllByAccountID(ctx context.Context, accountID string) ([]*entity.Transaction, error) {
	return nil, nil
}

func TestProcessTransactionRejectsInvalidType(t *testing.T) {
	acc, _ := entity.NewAccount("ACC-001", "USER-001")
	_ = acc.UpdateBalance(100)

	useCase := NewProcessTransactionUseCase(
		&accountRepoStub{account: acc},
		&txRepoStub{},
	)

	_, err := useCase.Execute(context.Background(), ProcessTransactionRequest{
		AccountID:   "ACC-001",
		Type:        "PIX",
		Amount:      10,
		Description: "teste",
	})

	if !errors.Is(err, ErrInvalidTransactionType) {
		t.Fatalf("expected ErrInvalidTransactionType, got %v", err)
	}
}

func TestProcessTransactionMapsAccountNotFound(t *testing.T) {
	useCase := NewProcessTransactionUseCase(
		&accountRepoStub{err: errors.New("db unavailable")},
		&txRepoStub{},
	)

	_, err := useCase.Execute(context.Background(), ProcessTransactionRequest{
		AccountID:   "ACC-404",
		Type:        "DEBIT",
		Amount:      10,
		Description: "teste",
	})

	if !errors.Is(err, ErrAccountNotFound) {
		t.Fatalf("expected ErrAccountNotFound, got %v", err)
	}
}

func TestProcessTransactionReturnsTransactionSaveError(t *testing.T) {
	acc, _ := entity.NewAccount("ACC-001", "USER-001")
	_ = acc.UpdateBalance(100)

	expectedErr := errors.New("failed to save transaction")
	useCase := NewProcessTransactionUseCase(
		&accountRepoStub{account: acc},
		&txRepoStub{err: expectedErr},
	)

	_, err := useCase.Execute(context.Background(), ProcessTransactionRequest{
		AccountID:   "ACC-001",
		Type:        "DEBIT",
		Amount:      10,
		Description: "teste",
	})

	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected transaction save error, got %v", err)
	}
}
