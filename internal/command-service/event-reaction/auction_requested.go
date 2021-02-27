package event_reaction

import (
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/adapter"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/command-service/port"
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
	sender := port.CreatePublisher(conf.RedisConf.WriteUrl, port.FakeLogger{}, port.BidChannel)
	return CreateAuctionRequestCommandWithInterfaces(handler, preserver, sender)
}

func CreateAuctionRequestCommandWithInterfaces(
	handler AuctionCreateEventHandler,
	preserver AuctionEventPreserver,
	sender Sender,
) AuctionRequestedCommand {
	return AuctionRequestedCommand{handler: handler, preserver: preserver, sender: sender}
}

func (a AuctionRequestedCommand) Execute(event domain.Event) {
	var auction domain.CreateAuctionRequested
	notifyEvent := domain.NotifyEvent{
		Event:         event.Event,
		CorrelationId: event.CorrelationId,
	}
	if err := json.Unmarshal(event.Payload, &auction); err != nil {
		notifyEvent.Error = fmt.Sprintf("can not unarshal event: %v", err)
		a.sender.NotifyUserFail(notifyEvent)
	}

	err := a.handler.Handle(auction)
	if err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with auction creating: %v", err)
		a.sender.NotifyUserFail(notifyEvent)
	}

	// TBD if it fails the firs Handle should also be reverted
	err = a.preserver.Handle(event.Event, auction)
	if err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with saving data: %v", err)
		a.sender.NotifyUserFail(notifyEvent)
	}

	a.sender.NotifyUserSuccess(notifyEvent)
	a.sender.PublishData(event)
}
