package command

import (
	event_reaction "github.com/lantosgyuri/auction-portal/internal/data-transformer/event-reaction"
	custom_error "github.com/lantosgyuri/auction-portal/internal/pkg/custom-error"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type DeleteBid struct {
	BidRepo BidRepository
}

func (d *DeleteBid) Execute(event event_reaction.Event) error {
	var bid event_reaction.BidDeletedEvent

	if err := marshal.Payload(event.Payload, &bid); err != nil {
		return err
	}

	if err := d.BidRepo.DeleteUserBid(bid); err != nil {
		return custom_error.Create("can not delete user bid", err)
	}

	return nil
}
