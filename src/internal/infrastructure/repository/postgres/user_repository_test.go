package postgres

import (
	"context"
	"testing"

	"github.com/victor-silveira/go-wallet-core/src/internal/domain/entity"
)

func TestUserRepositoryReturnsUserCopy(t *testing.T) {
	repo := NewUserRepository()
	user := &entity.User{
		ID:    "USER-001",
		Name:  "Victor",
		Email: "victor@teste.com",
	}

	if err := repo.Save(context.Background(), user); err != nil {
		t.Fatalf("save user failed: %v", err)
	}

	saved, err := repo.GetByID(context.Background(), "USER-001")
	if err != nil {
		t.Fatalf("get user failed: %v", err)
	}

	saved.Name = "Alterado"

	reloaded, err := repo.GetByID(context.Background(), "USER-001")
	if err != nil {
		t.Fatalf("reload user failed: %v", err)
	}

	if reloaded.Name != "Victor" {
		t.Fatalf("repository should return user copy, expected Victor got %s", reloaded.Name)
	}
}
