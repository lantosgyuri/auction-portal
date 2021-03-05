package command

import (
	event_reaction "github.com/lantosgyuri/auction-portal/internal/data-transformer/event-reaction"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type AuctionPreserver struct {
	AuctionRepo AuctionRepository
}

func (a *AuctionPreserver) Execute(event event_reaction.Event) error {
	var auction event_reaction.AuctionCreatedEvent

	if err := marshal.Payload(event.Payload, &auction); err != nil {
		return err
	}

	return a.AuctionRepo.SaveAuction(auction)
}
