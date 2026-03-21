package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/victor-silveira/go-wallet-core/src/interfaces/http/handler"
)

func TestHealthHandlerGET(t *testing.T) {
	h := handler.NewHealthHandler("9.9.9")
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	h.HealthCheck(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("code %d", rec.Code)
	}
	var out map[string]any
	if err := json.NewDecoder(rec.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}
	if out["status"] != "UP" || out["version"] != "9.9.9" {
		t.Fatalf("response %+v", out)
	}
}

func TestHealthHandlerMethodNotAllowed(t *testing.T) {
	h := handler.NewHealthHandler("1")
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	rec := httptest.NewRecorder()
	h.HealthCheck(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("code %d", rec.Code)
	}
}
