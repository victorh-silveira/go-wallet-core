package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/victor-silveira/go-wallet-core/src/application/user"
	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
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

	var req user.CreateUserRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		if errors.Is(err, io.EOF) {
			RespondWithError(w, http.StatusBadRequest, "Corpo da requisição está vazio")
			return
		}
		RespondWithError(w, http.StatusBadRequest, "Formato JSON inválido")
		return
	}

	res, err := h.createUserUseCase.Execute(r.Context(), req)
	if err != nil {
		if errors.Is(err, entity.ErrUserIDRequired) ||
			errors.Is(err, entity.ErrUserNameRequired) ||
			errors.Is(err, entity.ErrUserEmailRequired) {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		RespondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, res)
}
