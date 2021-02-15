package command

import "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"

type DeleteUserHandler struct {
	Repo UserRepository
}

func (d DeleteUserHandler) Handle(request domain.DeleteUserRequest) error {
	return d.Repo.DeleteUser(request)
}
