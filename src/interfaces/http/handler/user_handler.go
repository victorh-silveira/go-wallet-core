package handler

import (
	"encoding/json"
	"net/http"

	"github.com/victor-silveira/go-wallet-core/src/application/user"
)

type UserHandler struct {
	createUserUseCase *user.CreateUserUseCase
}

func NewUserHandler(createUserUseCase *user.CreateUserUseCase) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RespondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	if r.Body == nil || r.ContentLength == 0 {
		RespondWithError(w, http.StatusBadRequest, "Corpo da requisição está vazio")
		return
	}

	var req user.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Formato JSON inválido")
		return
	}

	res, err := h.createUserUseCase.Execute(r.Context(), req)
	if err != nil {
		RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, res)
}
