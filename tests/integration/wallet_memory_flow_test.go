package integration_test

import (
	"context"
	"testing"
	"time"

	appwallet "github.com/victor-silveira/go-wallet-core/src/application/wallet"
	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
	"github.com/victor-silveira/go-wallet-core/src/infrastructure/repository/memory"
)

func TestWalletMemoryFlowCreditAndListTransactions(t *testing.T) {
	repo := memory.NewWalletRepository()
	acc, err := entity.NewAccount("ACC-INT", "USER-INT")
	if err != nil {
		t.Fatal(err)
	}
	if err := acc.UpdateBalance(5000); err != nil {
		t.Fatal(err)
	}
	if err := repo.SaveAccount(context.Background(), acc); err != nil {
		t.Fatal(err)
	}

	uc := appwallet.NewProcessTransactionUseCase(repo, repo)
	res, err := uc.Execute(context.Background(), appwallet.ProcessTransactionRequest{
		AccountID: "ACC-INT", Type: "CREDIT", Amount: 100, Description: "int test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.NewBalance != 5100 {
		t.Fatalf("balance %d", res.NewBalance)
	}

	list, err := repo.FindAllByAccountID(context.Background(), "ACC-INT")
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 || list[0].Amount != 100 {
		t.Fatalf("transactions %+v", list)
	}
}

func TestWalletRepositoryGetByUserID(t *testing.T) {
	repo := memory.NewWalletRepository()
	acc := &entity.Account{
		ID: "A1", UserID: "UX", Balance: 1,
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}
	if err := repo.SaveAccount(context.Background(), acc); err != nil {
		t.Fatal(err)
	}
	got, err := repo.GetByUserID(context.Background(), "UX")
	if err != nil || got.ID != "A1" {
		t.Fatalf("got %+v err %v", got, err)
	}
}
