package command

import (
	"context"
	"errors"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

type CreateUserHandler struct {
	Repo UserRepository
}

func (c CreateUserHandler) Handle(context context.Context, userRequest domain.CreateUserRequested) error {

	if userRequest.Name == "" {
		return errors.New("no name provided for user")
	}

	if userRequest.Password == "" {
		return errors.New("no password is set")
	}

	user := domain.User{
		Name:     userRequest.Name,
		Password: userRequest.Password,
	}

	return c.Repo.CreateUser(user)
}
