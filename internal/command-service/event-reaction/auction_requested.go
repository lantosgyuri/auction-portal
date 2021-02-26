package event_reaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/adapter"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

type AuctionCreateEventHandler interface {
	Handle(event domain.CreateAuctionRequested) error
}
type AuctionEventPreserver interface {
	Handle(eventName string, event domain.AuctionEvent) error
}

type AuctionRequestedCommand struct {
	handler   AuctionCreateEventHandler
	preserver AuctionEventPreserver
}

func CreateAuctionRequestedCommand() AuctionRequestedCommand {
	handler := command.CreateAuctionHandler{
		Repo: adapter.CreateMariaDbAuctionRepository(),
	}
	preserver := command.SaveAuctionEventHandler{Repo: adapter.CreateMariaDbAuctionRepository()}

	return CreateAuctionRequestCommandWithInterfaces(handler, preserver)
}

func CreateAuctionRequestCommandWithInterfaces(handler AuctionCreateEventHandler, preserver AuctionEventPreserver) AuctionRequestedCommand {
	return AuctionRequestedCommand{handler: handler, preserver: preserver}
}

func (a AuctionRequestedCommand) Execute(event domain.Event) error {
	var auction domain.CreateAuctionRequested
	if err := json.Unmarshal(event.Payload, &auction); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling auction create: %v", err))
	}
	err := a.handler.Handle(auction)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened with creating auction: %v", err))
	}
	err = a.preserver.Handle(event.Event, auction)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened during saving the auction event: %v", err))
	}
	return nil
}
