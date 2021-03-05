package command

import (
	event_reaction "github.com/lantosgyuri/auction-portal/internal/data-transformer/event-reaction"
	custom_error "github.com/lantosgyuri/auction-portal/internal/pkg/custom-error"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type SaveWinner struct {
	AuctionRepo AuctionRepository
	WinnerRepo  WinnerRepository
}

func (s *SaveWinner) Execute(event event_reaction.Event) error {
	var winner event_reaction.WinnerAnnouncedEvent

	if err := marshal.Payload(event.Payload, &winner); err != nil {
		return err
	}

	if err := s.AuctionRepo.UpdateAuctionWinner(winner); err != nil {
		return custom_error.Create("can not update auction winner", err)
	}

	if err := s.WinnerRepo.SaveWinner(winner); err != nil {
		return custom_error.Create("can not save winner", err)
	}

	return nil
}
