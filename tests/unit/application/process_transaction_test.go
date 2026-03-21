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
	saved *entity.Transaction
	err   error
}

func (s *txRepoStub) SaveTransaction(ctx context.Context, transaction *entity.Transaction) error {
	if s.err != nil {
		return s.err
	}
	s.saved = transaction
	return nil
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

func TestProcessTransactionEmptyAccountID(t *testing.T) {
	uc := appwallet.NewProcessTransactionUseCase(&accountRepoStub{}, &txRepoStub{})
	_, err := uc.Execute(context.Background(), appwallet.ProcessTransactionRequest{
		AccountID: "", Type: "CREDIT", Amount: 1, Description: "x",
	})
	if !errors.Is(err, appwallet.ErrAccountNotFound) {
		t.Fatalf("got %v", err)
	}
}

func TestProcessTransactionInvalidAmount(t *testing.T) {
	acc, _ := entity.NewAccount("ACC-1", "U1")
	uc := appwallet.NewProcessTransactionUseCase(
		&accountRepoStub{account: acc},
		&txRepoStub{},
	)
	_, err := uc.Execute(context.Background(), appwallet.ProcessTransactionRequest{
		AccountID: "ACC-1", Type: "CREDIT", Amount: 0, Description: "x",
	})
	if !errors.Is(err, entity.ErrInvalidAmount) {
		t.Fatalf("got %v want ErrInvalidAmount", err)
	}
}

func TestProcessTransactionAccountNotFound(t *testing.T) {
	uc := appwallet.NewProcessTransactionUseCase(
		&accountRepoStub{err: errors.New("missing")},
		&txRepoStub{},
	)
	_, err := uc.Execute(context.Background(), appwallet.ProcessTransactionRequest{
		AccountID: "ACC-X", Type: "CREDIT", Amount: 10, Description: "x",
	})
	if !errors.Is(err, appwallet.ErrAccountNotFound) {
		t.Fatalf("got %v", err)
	}
}

func TestProcessTransactionCredit(t *testing.T) {
	acc, _ := entity.NewAccount("ACC-1", "U1")
	_ = acc.UpdateBalance(1000)
	repo := &accountRepoStub{account: acc}
	tx := &txRepoStub{}
	uc := appwallet.NewProcessTransactionUseCase(repo, tx)

	res, err := uc.Execute(context.Background(), appwallet.ProcessTransactionRequest{
		AccountID: "ACC-1", Type: "credit", Amount: 500, Description: "c",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.NewBalance != 1500 {
		t.Fatalf("balance got %d want 1500", res.NewBalance)
	}
	if repo.account == nil || repo.account.Balance != 1500 {
		t.Fatal("account not persisted correctly")
	}
	if tx.saved == nil || tx.saved.Type != entity.Credit || tx.saved.Amount != 500 {
		t.Fatalf("transaction: %+v", tx.saved)
	}
}

func TestProcessTransactionDebit(t *testing.T) {
	acc, _ := entity.NewAccount("ACC-1", "U1")
	_ = acc.UpdateBalance(1000)
	repo := &accountRepoStub{account: acc}
	tx := &txRepoStub{}
	uc := appwallet.NewProcessTransactionUseCase(repo, tx)

	res, err := uc.Execute(context.Background(), appwallet.ProcessTransactionRequest{
		AccountID: "ACC-1", Type: "DEBIT", Amount: 300, Description: "d",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.NewBalance != 700 {
		t.Fatalf("balance got %d want 700", res.NewBalance)
	}
	if tx.saved == nil || tx.saved.Type != entity.Debit {
		t.Fatalf("transaction: %+v", tx.saved)
	}
}

func TestProcessTransactionInsufficientBalance(t *testing.T) {
	acc, _ := entity.NewAccount("ACC-1", "U1")
	_ = acc.UpdateBalance(100)
	repo := &accountRepoStub{account: acc}
	uc := appwallet.NewProcessTransactionUseCase(repo, &txRepoStub{})

	_, err := uc.Execute(context.Background(), appwallet.ProcessTransactionRequest{
		AccountID: "ACC-1", Type: "DEBIT", Amount: 500, Description: "d",
	})
	if !errors.Is(err, entity.ErrInsufficientBalance) {
		t.Fatalf("got %v want ErrInsufficientBalance", err)
	}
}

func TestProcessTransactionSaveTransactionFails(t *testing.T) {
	acc, _ := entity.NewAccount("ACC-1", "U1")
	_ = acc.UpdateBalance(1000)
	repo := &accountRepoStub{account: acc}
	uc := appwallet.NewProcessTransactionUseCase(repo, &txRepoStub{err: errors.New("tx fail")})

	_, err := uc.Execute(context.Background(), appwallet.ProcessTransactionRequest{
		AccountID: "ACC-1", Type: "CREDIT", Amount: 10, Description: "c",
	})
	if err == nil || err.Error() != "tx fail" {
		t.Fatalf("got %v", err)
	}
}
