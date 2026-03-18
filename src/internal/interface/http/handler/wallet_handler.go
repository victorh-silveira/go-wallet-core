package handler

import (
	"encoding/json"
	"net/http"

	"github.com/victor-silveira/go-wallet-core/src/internal/usecase/wallet"
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
		RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, res)
}
