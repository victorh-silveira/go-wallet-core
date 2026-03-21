package entity_test

import (
	"errors"
	"testing"

	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
)

func TestNewUserValidation(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		nm      string
		email   string
		wantErr error
	}{
		{"empty id", "", "n", "e@x", entity.ErrUserIDRequired},
		{"empty name", "1", "", "e@x", entity.ErrUserNameRequired},
		{"empty email", "1", "n", "", entity.ErrUserEmailRequired},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := entity.NewUser(tc.id, tc.nm, tc.email)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("got %v want %v", err, tc.wantErr)
			}
		})
	}
}

func TestNewUserOK(t *testing.T) {
	u, err := entity.NewUser("USER-1", "Nome", "a@b.c")
	if err != nil {
		t.Fatal(err)
	}
	if u.ID != "USER-1" || u.Name != "Nome" || u.Email != "a@b.c" {
		t.Fatalf("unexpected user: %+v", u)
	}
}
