package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type SaveAuctionEventHandler struct {
	Repo AuctionRepository
}

func (s SaveAuctionEventHandler) Handle(event domain.Event) error {
	eventType := event.Event

	auctionEvent := domain.AuctionEvent{}

	switch eventType {
	case domain.AuctionRequested:
		var auctionCreate domain.CreateAuctionMessage
		if err := json.Unmarshal(event.Payload, &auctionCreate); err != nil {
			return errors.New("can not marshal event payload")
		}
		auctionEvent.Name = auctionCreate.Name
		auctionEvent.DueDate = auctionCreate.DueDate
		auctionEvent.StartDate = auctionCreate.StartDate
		auctionEvent.Timestamp = auctionCreate.Timestamp

	case domain.AuctionWinnerAnnounced:
		var winner domain.AuctionWinnerMessage
		if err := json.Unmarshal(event.Payload, &winner); err != nil {
			return errors.New("can not marshal event payload")
		}
		auctionEvent.AuctionId = winner.AuctionId
		auctionEvent.Winner = winner.WinnerId
		auctionEvent.Timestamp = winner.Timestamp

	default:
		return errors.New(fmt.Sprintf("no event found for: %v", eventType))
	}

	return s.Repo.SaveAuctionEvent(auctionEvent)
}
