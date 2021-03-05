package command

import (
	event_reaction "github.com/lantosgyuri/auction-portal/internal/data-transformer/event-reaction"
	custom_error "github.com/lantosgyuri/auction-portal/internal/pkg/custom-error"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type DeleteUser struct {
	UserRepo UserRepository
	BidRepo  BidRepository
}

func (d *DeleteUser) Execute(event event_reaction.Event) error {
	var user event_reaction.UserDeletedEvent

	if err := marshal.Payload(event.Payload, &user); err != nil {
		return err
	}

	if err := d.UserRepo.DeleteUser(user); err != nil {
		return custom_error.Create("can not delete user", err)
	}

	if err := d.BidRepo.DeleteUserEntries(user); err != nil {
		return custom_error.Create("can not delete user bid entries", err)
	}

	return nil
}
