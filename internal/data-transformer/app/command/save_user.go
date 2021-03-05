package command

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type SaveUser struct {
	Repo UserRepository
}

func (s *SaveUser) Execute(event domain.Event) error {
	var user domain.CreateUserRequested

	if err := marshal.Payload(event, &user); err != nil {
		return err
	}

	return s.Repo.SaveUser(user)
}
