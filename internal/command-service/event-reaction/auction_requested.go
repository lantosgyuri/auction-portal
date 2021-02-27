package event_reaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/adapter"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
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
	sender    Sender
}

func CreateAuctionRequestedCommand(conf config.CommandService) AuctionRequestedCommand {
	handler := command.CreateAuctionHandler{
		Repo: adapter.CreateMariaDbAuctionRepository(),
	}
	preserver := command.SaveAuctionEventHandler{Repo: adapter.CreateMariaDbAuctionRepository()}

	return CreateAuctionRequestCommandWithInterfaces(handler, preserver)
}

func CreateAuctionRequestCommandWithInterfaces(handler AuctionCreateEventHandler, preserver AuctionEventPreserver) AuctionRequestedCommand {
	return AuctionRequestedCommand{handler: handler, preserver: preserver}
}

func (a AuctionRequestedCommand) Execute(event domain.Event) {
	var auction domain.CreateAuctionRequested
	if err := json.Unmarshal(event.Payload, &auction); err != nil {
		a.sender.NotifyUserFail(
			event.CorrelationId,
			event.Event,
			errors.New(fmt.Sprintf("can not unarshal event: %v", err)))
	}
	err := a.handler.Handle(auction)
	if err != nil {
		a.sender.NotifyUserFail(
			event.CorrelationId,
			event.Event,
			errors.New(fmt.Sprintf("error happened with auction creating: %v", err)))
	}

	// TBD if it fails the firs Handle should also be reverted
	err = a.preserver.Handle(event.Event, auction)
	if err != nil {
		a.sender.NotifyUserFail(
			event.CorrelationId,
			event.Event,
			errors.New(fmt.Sprintf("error happened with saving data: %v", err)))
	}

	a.sender.NotifyUserSuccess(event.CorrelationId, event.Event)
	a.sender.PublishData(event)
}
