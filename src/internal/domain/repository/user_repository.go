package repository

import (
	"context"

	"github.com/victor-silveira/go-wallet-core/src/internal/domain/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
}
