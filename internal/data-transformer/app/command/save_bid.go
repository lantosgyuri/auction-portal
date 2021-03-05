package command

import (
	event_reaction "github.com/lantosgyuri/auction-portal/internal/data-transformer/event-reaction"
	custom_error "github.com/lantosgyuri/auction-portal/internal/pkg/custom-error"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type SaveBid struct {
	BidRepo     BidRepository
	AuctionRepo AuctionRepository
}

func (s *SaveBid) Execute(event event_reaction.Event) error {
	var bid event_reaction.BidPlacedEvent

	if err := marshal.Payload(event.Payload, &bid); err != nil {
		return err
	}

	if err := s.AuctionRepo.UpdateAuctionBid(bid); err != nil {
		return custom_error.Create("can not update auction", err)
	}

	if !bid.Rollback {
		if err := s.BidRepo.SaveUserBid(bid); err != nil {
			return custom_error.Create("can not save users bid", err)
		}
	}

	return nil
}
