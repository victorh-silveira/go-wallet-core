package wallet

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
	"github.com/victor-silveira/go-wallet-core/src/domain/repository"
)

type ProcessTransactionRequest struct {
	AccountID   string `json:"account_id"`
	Type        string `json:"type"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

type ProcessTransactionResponse struct {
	NewBalance int64 `json:"new_balance"`
}

type ProcessTransactionUseCase struct {
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
}

var (
	ErrInvalidTransactionType = errors.New("invalid transaction type")
	ErrAccountNotFound        = errors.New("account not found")
)

func NewProcessTransactionUseCase(accountRepo repository.AccountRepository, transactionRepo repository.TransactionRepository) *ProcessTransactionUseCase {
	return &ProcessTransactionUseCase{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (u *ProcessTransactionUseCase) Execute(ctx context.Context, request ProcessTransactionRequest) (ProcessTransactionResponse, error) {
	if request.AccountID == "" {
		return ProcessTransactionResponse{}, ErrAccountNotFound
	}

	if request.Amount <= 0 {
		return ProcessTransactionResponse{}, entity.ErrInvalidAmount
	}

	acc, err := u.accountRepo.GetByID(ctx, request.AccountID)
	if err != nil {
		return ProcessTransactionResponse{}, ErrAccountNotFound
	}

	var tType entity.TransactionType
	var signedDelta int64
	requestType := strings.ToUpper(strings.TrimSpace(request.Type))

	if requestType == string(entity.Debit) {
		tType = entity.Debit
		signedDelta = -request.Amount
	} else if requestType == string(entity.Credit) {
		tType = entity.Credit
		signedDelta = request.Amount
	} else {
		return ProcessTransactionResponse{}, ErrInvalidTransactionType
	}

	if err := acc.UpdateBalance(signedDelta); err != nil {
		return ProcessTransactionResponse{}, err
	}

	if err := u.accountRepo.SaveAccount(ctx, acc); err != nil {
		return ProcessTransactionResponse{}, err
	}

	id, err := newTransactionID()
	if err != nil {
		return ProcessTransactionResponse{}, err
	}
	trx, err := entity.NewTransaction(id, acc.ID, request.Description, tType, request.Amount)
	if err != nil {
		return ProcessTransactionResponse{}, err
	}
	if err := u.transactionRepo.SaveTransaction(ctx, trx); err != nil {
		return ProcessTransactionResponse{}, err
	}

	return ProcessTransactionResponse{
		NewBalance: acc.Balance,
	}, nil
}

func newTransactionID() (string, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("gerar id de transacao: %w", err)
	}
	return "TRX-" + hex.EncodeToString(b), nil
}
