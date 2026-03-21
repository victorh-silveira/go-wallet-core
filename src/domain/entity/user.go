package entity

import "errors"

var (
	ErrUserIDRequired    = errors.New("id is required")
	ErrUserNameRequired  = errors.New("name is required")
	ErrUserEmailRequired = errors.New("email is required")
)

type User struct {
	ID    string
	Name  string
	Email string
}

func NewUser(id, name, email string) (*User, error) {
	if id == "" {
		return nil, ErrUserIDRequired
	}
	if name == "" {
		return nil, ErrUserNameRequired
	}
	if email == "" {
		return nil, ErrUserEmailRequired
	}
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}, nil
}
