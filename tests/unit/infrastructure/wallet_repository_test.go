package infrastructure_test

import (
	"context"
	"testing"
	"time"

	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
	"github.com/victor-silveira/go-wallet-core/src/infrastructure/repository/memory"
)

func TestWalletRepositoryReturnsAccountCopy(t *testing.T) {
	repo := memory.NewWalletRepository()
	account := &entity.Account{
		ID:        "ACC-001",
		UserID:    "USER-001",
		Balance:   5000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repo.SaveAccount(context.Background(), account); err != nil {
		t.Fatalf("save account failed: %v", err)
	}

	saved, err := repo.GetByID(context.Background(), "ACC-001")
	if err != nil {
		t.Fatalf("get account failed: %v", err)
	}

	saved.Balance = 999

	reloaded, err := repo.GetByID(context.Background(), "ACC-001")
	if err != nil {
		t.Fatalf("reload account failed: %v", err)
	}

	if reloaded.Balance != 5000 {
		t.Fatalf("repository should be immutable outside, expected 5000 got %v", reloaded.Balance)
	}
}

func TestWalletRepositorySaveTransactionAndFind(t *testing.T) {
	repo := memory.NewWalletRepository()
	trx := &entity.Transaction{
		ID: "T1", AccountID: "ACC-001", Type: entity.Credit, Amount: 10,
		Description: "x", CreatedAt: time.Now(),
	}
	if err := repo.SaveTransaction(context.Background(), trx); err != nil {
		t.Fatal(err)
	}
	list, err := repo.FindAllByAccountID(context.Background(), "ACC-001")
	if err != nil || len(list) != 1 || list[0].ID != "T1" {
		t.Fatalf("list %+v err %v", list, err)
	}
}

func TestWalletRepositoryGetByIDNotFound(t *testing.T) {
	repo := memory.NewWalletRepository()
	_, err := repo.GetByID(context.Background(), "nope")
	if err == nil {
		t.Fatal("expected error")
	}
}
