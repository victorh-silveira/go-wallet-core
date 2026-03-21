package user

import (
	"context"

	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
	"github.com/victor-silveira/go-wallet-core/src/domain/repository"
)

type CreateUserRequest struct {
	ID    string
	Name  string
	Email string
}

type CreateUserResponse struct {
	ID    string
	Name  string
	Email string
}

type CreateUserUseCase struct {
	userRepo repository.UserRepository
}

func NewCreateUserUseCase(userRepo repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepo: userRepo,
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, request CreateUserRequest) (CreateUserResponse, error) {
	domainUser, err := entity.NewUser(request.ID, request.Name, request.Email)
	if err != nil {
		return CreateUserResponse{}, err
	}

	err = u.userRepo.Save(ctx, domainUser)
	if err != nil {
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{
		ID:    domainUser.ID,
		Name:  domainUser.Name,
		Email: domainUser.Email,
	}, nil
}
