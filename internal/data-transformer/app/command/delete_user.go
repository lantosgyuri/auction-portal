package command

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type DeleteUser struct {
	Repo UserRepository
}

func (d *DeleteUser) Execute(event domain.Event) error {
	var user domain.DeleteUserRequest

	if err := marshal.Payload(event, &user); err != nil {
		return err
	}

	return d.Repo.DeleteUser(user)
}
