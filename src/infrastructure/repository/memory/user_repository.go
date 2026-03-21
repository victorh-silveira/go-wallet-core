package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
)

type UserRepository struct {
	mu    sync.RWMutex
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
	if user == nil {
		return errors.New("user is required")
	}
	userCopy := *user
	r.users[user.ID] = &userCopy
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	userCopy := *user
	return &userCopy, nil
}
