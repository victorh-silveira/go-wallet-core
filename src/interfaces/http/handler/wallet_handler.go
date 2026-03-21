package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/victor-silveira/go-wallet-core/src/application/wallet"
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
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Corpo da requisição inválido")
		return
	}

	res, err := h.processTrxUseCase.Execute(r.Context(), req)
	if err != nil {
		if errors.Is(err, wallet.ErrInvalidTransactionType) {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, wallet.ErrAccountNotFound) {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, res)
}
