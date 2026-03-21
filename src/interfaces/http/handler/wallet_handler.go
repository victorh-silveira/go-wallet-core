package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/victor-silveira/go-wallet-core/src/application/wallet"
	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
)

type WalletHandler struct {
	processTrxUseCase *wallet.ProcessTransactionUseCase
}

func NewWalletHandler(processTrxUseCase *wallet.ProcessTransactionUseCase) *WalletHandler {
	return &WalletHandler{
		processTrxUseCase: processTrxUseCase,
	}
}

func (h *WalletHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RespondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	var req wallet.ProcessTransactionRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Corpo da requisição inválido")
		return
	}

	res, err := h.processTrxUseCase.Execute(r.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, wallet.ErrInvalidTransactionType):
			RespondWithError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, wallet.ErrAccountNotFound):
			RespondWithError(w, http.StatusNotFound, err.Error())
		case errors.Is(err, entity.ErrInvalidAmount):
			RespondWithError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, entity.ErrInsufficientBalance):
			RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		default:
			RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusOK, res)
}
