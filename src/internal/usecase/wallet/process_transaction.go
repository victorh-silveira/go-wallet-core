package wallet

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/victor-silveira/go-wallet-core/src/internal/domain/entity"
	"github.com/victor-silveira/go-wallet-core/src/internal/domain/repository"
)

type ProcessTransactionRequest struct {
	AccountID   string  `json:"account_id"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

type ProcessTransactionResponse struct {
	NewBalance float64 `json:"new_balance"`
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

	acc, err := u.accountRepo.GetByID(ctx, request.AccountID)
	if err != nil {
		return ProcessTransactionResponse{}, ErrAccountNotFound
	}

	var tType entity.TransactionType
	signedAmount := request.Amount
	requestType := strings.ToUpper(strings.TrimSpace(request.Type))

	if requestType == string(entity.Debit) {
		tType = entity.Debit
		signedAmount = -request.Amount
	} else if requestType == string(entity.Credit) {
		tType = entity.Credit
	} else {
		return ProcessTransactionResponse{}, ErrInvalidTransactionType
	}

	if err := acc.UpdateBalance(signedAmount); err != nil {
		return ProcessTransactionResponse{}, err
	}

	if err := u.accountRepo.SaveAccount(ctx, acc); err != nil {
		return ProcessTransactionResponse{}, err
	}

	id := fmt.Sprintf("TRX-%d", time.Now().UnixNano())
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
