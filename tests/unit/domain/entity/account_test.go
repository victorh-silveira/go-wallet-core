package entity_test

import (
	"testing"

	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
)

func TestNewAccountRequiresID(t *testing.T) {
	_, err := entity.NewAccount("", "USER-1")
	if err == nil {
		t.Fatal("expected error for empty account id")
	}
}

func TestNewAccountRequiresUserID(t *testing.T) {
	_, err := entity.NewAccount("ACC-1", "")
	if err == nil {
		t.Fatal("expected error for empty user id")
	}
}

func TestNewAccountStartsWithZeroBalance(t *testing.T) {
	a, err := entity.NewAccount("ACC-1", "USER-1")
	if err != nil {
		t.Fatal(err)
	}
	if a.Balance != 0 {
		t.Fatalf("balance: got %d want 0", a.Balance)
	}
}

func TestUpdateBalanceCredit(t *testing.T) {
	a, _ := entity.NewAccount("ACC-1", "USER-1")
	if err := a.UpdateBalance(100); err != nil {
		t.Fatal(err)
	}
	if a.Balance != 100 {
		t.Fatalf("balance: got %d want 100", a.Balance)
	}
}

func TestUpdateBalanceDebit(t *testing.T) {
	a, _ := entity.NewAccount("ACC-1", "USER-1")
	_ = a.UpdateBalance(500)
	if err := a.UpdateBalance(-200); err != nil {
		t.Fatal(err)
	}
	if a.Balance != 300 {
		t.Fatalf("balance: got %d want 300", a.Balance)
	}
}

func TestUpdateBalanceInsufficient(t *testing.T) {
	a, _ := entity.NewAccount("ACC-1", "USER-1")
	_ = a.UpdateBalance(100)
	err := a.UpdateBalance(-200)
	if err == nil {
		t.Fatal("expected ErrInsufficientBalance")
	}
	if err != entity.ErrInsufficientBalance {
		t.Fatalf("got %v want %v", err, entity.ErrInsufficientBalance)
	}
}
