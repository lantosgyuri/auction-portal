package command

import (
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type SaveBidEventHandler struct {
	Repo BidRepository
}

func (s SaveBidEventHandler) Handle(eventName string, event domain.BidEvent) error {
	rawEvent := domain.BidEventRaw{
		EventType: eventName,
	}

	switch e := event.(type) {
	case domain.BidPlaced:
		rawEvent.AuctionId = e.AuctionId
		rawEvent.Amount = e.Amount
		rawEvent.UserId = e.UserId
		rawEvent.TimeStamp = e.TimeStamp
		return s.Repo.SaveBidEvent(rawEvent)
	case domain.BidDeleted:
		rawEvent.AuctionId = e.AuctionId
		rawEvent.UserId = e.UserId
		rawEvent.BidId = e.BidId
		rawEvent.TimeStamp = e.TimeStamp
		return s.Repo.SaveBidEvent(rawEvent)
	default:
		return errors.New(fmt.Sprintf("no event found for: %v", e))
	}
}
