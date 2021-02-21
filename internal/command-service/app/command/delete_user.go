package command

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

type DeleteUserHandler struct {
	Repo UserRepository
}

func (d DeleteUserHandler) Handle(context context.Context, request domain.DeleteUserRequest) error {
	return d.Repo.DeleteUser(domain.User{
		Name: request.Name,
		Id:   request.Id,
	})
}
