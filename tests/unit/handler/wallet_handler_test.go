package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	appwallet "github.com/victor-silveira/go-wallet-core/src/application/wallet"
	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
	"github.com/victor-silveira/go-wallet-core/src/infrastructure/repository/memory"
	"github.com/victor-silveira/go-wallet-core/src/interfaces/http/handler"
)

func seedWallet(t *testing.T) *memory.WalletRepository {
	t.Helper()
	repo := memory.NewWalletRepository()
	acc := &entity.Account{
		ID:        "ACC-H1",
		UserID:    "U1",
		Balance:   10_000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := repo.SaveAccount(context.Background(), acc); err != nil {
		t.Fatal(err)
	}
	return repo
}

func TestWalletHandlerMethodNotAllowed(t *testing.T) {
	repo := seedWallet(t)
	h := handler.NewWalletHandler(appwallet.NewProcessTransactionUseCase(repo, repo))
	req := httptest.NewRequest(http.MethodGet, "/wallet/transaction", nil)
	rec := httptest.NewRecorder()
	h.Transaction(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("code %d", rec.Code)
	}
}

func TestWalletHandlerCredit(t *testing.T) {
	repo := seedWallet(t)
	h := handler.NewWalletHandler(appwallet.NewProcessTransactionUseCase(repo, repo))
	body := `{"account_id":"ACC-H1","type":"CREDIT","amount":100,"description":"x"}`
	req := httptest.NewRequest(http.MethodPost, "/wallet/transaction", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.Transaction(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("code %d %s", rec.Code, rec.Body.String())
	}
	var out struct {
		NewBalance int64 `json:"new_balance"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}
	if out.NewBalance != 10100 {
		t.Fatalf("balance %d want 10100", out.NewBalance)
	}
}

func TestWalletHandlerNotFound(t *testing.T) {
	repo := seedWallet(t)
	h := handler.NewWalletHandler(appwallet.NewProcessTransactionUseCase(repo, repo))
	body := `{"account_id":"UNKNOWN","type":"CREDIT","amount":1,"description":"x"}`
	req := httptest.NewRequest(http.MethodPost, "/wallet/transaction", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.Transaction(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("code %d", rec.Code)
	}
}

func TestWalletHandlerInsufficientBalance(t *testing.T) {
	repo := seedWallet(t)
	h := handler.NewWalletHandler(appwallet.NewProcessTransactionUseCase(repo, repo))
	body := `{"account_id":"ACC-H1","type":"DEBIT","amount":999999,"description":"x"}`
	req := httptest.NewRequest(http.MethodPost, "/wallet/transaction", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.Transaction(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("code %d %s", rec.Code, rec.Body.String())
	}
}
