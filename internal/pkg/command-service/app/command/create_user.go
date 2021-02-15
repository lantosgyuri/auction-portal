package command

import (
	"errors"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type CreateUserHandler struct {
	Repo UserRepository
}

func (c CreateUserHandler) Handle(userRequest domain.CreateUserRequested) error {

	if userRequest.Name == "" {
		return errors.New("no name provided for user")
	}

	if userRequest.Password == "" {
		return errors.New("no password is set")
	}

	return c.Repo.CreateUser(userRequest)
}
