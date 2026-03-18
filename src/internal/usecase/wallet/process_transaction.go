package wallet

import (
	"context"
	"fmt"
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

func NewProcessTransactionUseCase(accountRepo repository.AccountRepository, transactionRepo repository.TransactionRepository) *ProcessTransactionUseCase {
	return &ProcessTransactionUseCase{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (u *ProcessTransactionUseCase) Execute(ctx context.Context, request ProcessTransactionRequest) (ProcessTransactionResponse, error) {
	acc, err := u.accountRepo.GetByID(ctx, request.AccountID)
	if err != nil {
		return ProcessTransactionResponse{}, err
	}

	var tType entity.TransactionType
	signedAmount := request.Amount

	if request.Type == "DEBIT" {
		tType = entity.Debit
		signedAmount = -request.Amount
	} else {
		tType = entity.Credit
	}

	if err := acc.UpdateBalance(signedAmount); err != nil {
		return ProcessTransactionResponse{}, err
	}

	if err := u.accountRepo.SaveAccount(ctx, acc); err != nil {
		return ProcessTransactionResponse{}, err
	}

	id := fmt.Sprintf("TRX-%d", time.Now().UnixNano())
	trx, _ := entity.NewTransaction(id, acc.ID, request.Description, tType, request.Amount)
	_ = u.transactionRepo.SaveTransaction(ctx, trx)

	return ProcessTransactionResponse{
		NewBalance: acc.Balance,
	}, nil
}
