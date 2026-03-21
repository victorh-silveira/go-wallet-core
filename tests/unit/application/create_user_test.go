package application_test

import (
	"context"
	"errors"
	"testing"

	appuser "github.com/victor-silveira/go-wallet-core/src/application/user"
	"github.com/victor-silveira/go-wallet-core/src/domain/entity"
)

type userRepoStub struct {
	saveErr error
	saved   *entity.User
}

func (s *userRepoStub) Save(ctx context.Context, user *entity.User) error {
	if s.saveErr != nil {
		return s.saveErr
	}
	s.saved = user
	return nil
}

func (s *userRepoStub) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return nil, errors.New("not implemented")
}

func TestCreateUserSuccess(t *testing.T) {
	repo := &userRepoStub{}
	uc := appuser.NewCreateUserUseCase(repo)
	res, err := uc.Execute(context.Background(), appuser.CreateUserRequest{
		ID: "U1", Name: "N", Email: "n@e.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if res.ID != "U1" || repo.saved == nil || repo.saved.ID != "U1" {
		t.Fatalf("unexpected result: %+v repo.saved=%v", res, repo.saved)
	}
}

func TestCreateUserValidation(t *testing.T) {
	uc := appuser.NewCreateUserUseCase(&userRepoStub{})
	_, err := uc.Execute(context.Background(), appuser.CreateUserRequest{
		ID: "", Name: "N", Email: "n@e.com",
	})
	if !errors.Is(err, entity.ErrUserIDRequired) {
		t.Fatalf("got %v", err)
	}
}

func TestCreateUserSaveError(t *testing.T) {
	repo := &userRepoStub{saveErr: errors.New("db down")}
	uc := appuser.NewCreateUserUseCase(repo)
	_, err := uc.Execute(context.Background(), appuser.CreateUserRequest{
		ID: "U1", Name: "N", Email: "n@e.com",
	})
	if err == nil || err.Error() != "db down" {
		t.Fatalf("got %v", err)
	}
}
