package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/victor-silveira/go-wallet-core/src/internal/domain/entity"
)

func TestWalletRepositoryReturnsAccountCopy(t *testing.T) {
	repo := NewWalletRepository()
	account := &entity.Account{
		ID:        "ACC-001",
		UserID:    "USER-001",
		Balance:   50,
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

	if reloaded.Balance != 50 {
		t.Fatalf("repository should be immutable outside, expected 50 got %v", reloaded.Balance)
	}
}

func TestWalletRepositoryReturnsTransactionCopies(t *testing.T) {
	repo := NewWalletRepository()
	trx, _ := entity.NewTransaction("TRX-001", "ACC-001", "teste", entity.Credit, 10)

	if err := repo.SaveTransaction(context.Background(), trx); err != nil {
		t.Fatalf("save transaction failed: %v", err)
	}

	list, err := repo.FindAllByAccountID(context.Background(), "ACC-001")
	if err != nil {
		t.Fatalf("list transaction failed: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 transaction, got %d", len(list))
	}

	list[0].Amount = 999

	listAgain, err := repo.FindAllByAccountID(context.Background(), "ACC-001")
	if err != nil {
		t.Fatalf("list transaction failed: %v", err)
	}

	if listAgain[0].Amount != 10 {
		t.Fatalf("repository should return copies, expected 10 got %v", listAgain[0].Amount)
	}
}
