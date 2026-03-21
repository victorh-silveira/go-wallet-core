package application_test

import (
	"context"
	"errors"
	"testing"

	appwallet "github.com/victor-silveira/go-wallet-core/src/application/wallet"
	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
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

func (s *accountRepoStub) UpdateBalance(ctx context.Context, accountID string, deltaCentavos int64) error {
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
	_ = acc.UpdateBalance(10_000)

	useCase := appwallet.NewProcessTransactionUseCase(
		&accountRepoStub{account: acc},
		&txRepoStub{},
	)

	_, err := useCase.Execute(context.Background(), appwallet.ProcessTransactionRequest{
		AccountID:   "ACC-001",
		Type:        "PIX",
		Amount:      10,
		Description: "teste",
	})

	if !errors.Is(err, appwallet.ErrInvalidTransactionType) {
		t.Fatalf("expected ErrInvalidTransactionType, got %v", err)
	}
}
