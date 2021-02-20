package event_reaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

func init() {
	Commands[domain.AuctionRequested] = AuctionRequestedCommand{}
}

type AuctionRequestedCommand struct{}

func (a AuctionRequestedCommand) Execute(application app.Application, event domain.Event) error {
	var auction domain.CreateAuctionRequested
	if err := json.Unmarshal(event.Payload, &auction); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling auction create: %v", err))
	}
	err := application.Commands.CreateAuction.Handle(auction)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened with creating auction: %v", err))
	}
	err = application.Commands.SaveAuctionEvent.Handle(event.Event, auction)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened during saving the auction event: %v", err))
	}

	return nil
}
