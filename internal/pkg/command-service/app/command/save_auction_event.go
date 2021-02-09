package command

import (
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type SaveAuctionEventHandler struct {
	Repo Repository
}

func (s SaveAuctionEventHandler) Handle(event domain.Event) error {
	nEvent := domain.NormalizedAuctionEvent{
		Event: event.Event,
		Data:  string(event.Payload),
	}
	return s.Repo.SaveAuctionEvent(nEvent)
}
