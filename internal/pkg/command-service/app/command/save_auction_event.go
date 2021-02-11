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

	switch eventType {
	case domain.AuctionRequested:
		var auctionCreate domain.CreateAuctionRequested
		if err := json.Unmarshal(event.Payload, &auctionCreate); err != nil {
			return errors.New("can not marshal event payload")
		}
		return s.Repo.SaveAuctionEvent(auctionCreate)
	case domain.AuctionWinnerAnnounced:
		var winner domain.WinnerAnnounced
		if err := json.Unmarshal(event.Payload, &winner); err != nil {
			return errors.New("can not marshal event payload")
		}
		return s.Repo.SaveAuctionEvent(winner)
	default:
		return errors.New(fmt.Sprintf("no event found for: %v", eventType))
	}
}
