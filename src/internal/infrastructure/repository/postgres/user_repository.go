package postgres

import (
	"context"
	"errors"
	"sync"

	"github.com/victor-silveira/go-wallet-core/src/internal/domain/entity"
)

type UserRepository struct {
	mu    sync.Mutex
	users map[string]*entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*entity.User),
	}
}

func (r *UserRepository) Save(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
