package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	appuser "github.com/victor-silveira/go-wallet-core/src/application/user"
	"github.com/victor-silveira/go-wallet-core/src/infrastructure/repository/memory"
	"github.com/victor-silveira/go-wallet-core/src/interfaces/http/handler"
)

func TestUserHandlerCreateUserMethodNotAllowed(t *testing.T) {
	h := handler.NewUserHandler(appuser.NewCreateUserUseCase(memory.NewUserRepository()))
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	h.CreateUser(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("code %d", rec.Code)
	}
}

func TestUserHandlerCreateUserEmptyBody(t *testing.T) {
	h := handler.NewUserHandler(appuser.NewCreateUserUseCase(memory.NewUserRepository()))
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(""))
	rec := httptest.NewRecorder()
	h.CreateUser(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("code %d", rec.Code)
	}
}

func TestUserHandlerCreateUserInvalidJSON(t *testing.T) {
	h := handler.NewUserHandler(appuser.NewCreateUserUseCase(memory.NewUserRepository()))
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.CreateUser(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("code %d", rec.Code)
	}
}

func TestUserHandlerCreateUserUnknownField(t *testing.T) {
	h := handler.NewUserHandler(appuser.NewCreateUserUseCase(memory.NewUserRepository()))
	body := `{"id":"U1","name":"N","email":"a@b.c","extra":true}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.CreateUser(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("code %d", rec.Code)
	}
}

func TestUserHandlerCreateUserSuccess(t *testing.T) {
	repo := memory.NewUserRepository()
	h := handler.NewUserHandler(appuser.NewCreateUserUseCase(repo))
	body := `{"id":"user-h1","name":"Test","email":"t@test.com"}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.CreateUser(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("code %d body %s", rec.Code, rec.Body.String())
	}
	var out map[string]any
	if err := json.NewDecoder(rec.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}
	if out["id"] != "user-h1" {
		t.Fatalf("response %+v", out)
	}
}

func TestUserHandlerCreateUserValidation(t *testing.T) {
	h := handler.NewUserHandler(appuser.NewCreateUserUseCase(memory.NewUserRepository()))
	body := `{"id":"","name":"N","email":"a@b.c"}`
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.CreateUser(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("code %d", rec.Code)
	}
}
