package command

import (
	event_reaction "github.com/lantosgyuri/auction-portal/internal/data-transformer/event-reaction"
	custom_error "github.com/lantosgyuri/auction-portal/internal/pkg/custom-error"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type SaveUser struct {
	UserRepo UserRepository
	BidRepo  BidRepository
}

func (s *SaveUser) Execute(event event_reaction.Event) error {
	var user event_reaction.UserCreatedEvent

	if err := marshal.Payload(event.Payload, &user); err != nil {
		return err
	}

	if err := s.UserRepo.SaveUser(user); err != nil {
		return custom_error.Create("can not save user", err)
	}

	if err := s.BidRepo.CreateUser(user); err != nil {
		return custom_error.Create("can not create entries for user", err)
	}

	return nil
}
