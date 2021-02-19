package command

import (
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type SaveAuctionEventHandler struct {
	Repo AuctionRepository
}

func (s SaveAuctionEventHandler) Handle(eventName string, event domain.AuctionEvent) error {
	rawEvent := domain.AuctionEventRaw{
		EventType: eventName,
	}
	switch e := event.(type) {
	case domain.CreateAuctionRequested:
		rawEvent.Name = e.Name
		rawEvent.DueDate = e.DueDate
		rawEvent.StartDate = e.StartDate
		return s.Repo.SaveAuctionEvent(rawEvent)
	case domain.WinnerAnnounced:
		rawEvent.Winner = e.WinnerId
		return s.Repo.SaveAuctionEvent(rawEvent)
	default:
		return errors.New(fmt.Sprintf("no event found for: %v", e))
	}
}
