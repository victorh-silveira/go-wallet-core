package entity_test

import (
	"errors"
	"testing"

	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
)

func TestNewTransactionRequiresIDs(t *testing.T) {
	_, err := entity.NewTransaction("", "ACC-1", "d", entity.Credit, 100)
	if err == nil {
		t.Fatal("expected error for empty transaction id")
	}
	_, err = entity.NewTransaction("TRX-1", "", "d", entity.Credit, 100)
	if err == nil {
		t.Fatal("expected error for empty account id")
	}
}

func TestNewTransactionRequiresPositiveAmount(t *testing.T) {
	_, err := entity.NewTransaction("TRX-1", "ACC-1", "d", entity.Credit, 0)
	if !errors.Is(err, entity.ErrInvalidAmount) {
		t.Fatalf("got %v want %v", err, entity.ErrInvalidAmount)
	}
	_, err = entity.NewTransaction("TRX-1", "ACC-1", "d", entity.Debit, -1)
	if !errors.Is(err, entity.ErrInvalidAmount) {
		t.Fatalf("got %v want %v", err, entity.ErrInvalidAmount)
	}
}

func TestNewTransactionOK(t *testing.T) {
	trx, err := entity.NewTransaction("TRX-1", "ACC-1", "desc", entity.Credit, 50)
	if err != nil {
		t.Fatal(err)
	}
	if trx.Amount != 50 || trx.Type != entity.Credit {
		t.Fatalf("unexpected trx: %+v", trx)
	}
}
