package event_reaction

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/adapter"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/command-service/port"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
)

type BidPLacedEventHandler interface {
	Handle(ctx context.Context, placed domain.BidPlaced) error
}

type PreserveBidEvent interface {
	Handle(eventName string, event domain.BidEvent) error
}

type BidPlaceRequestedCommand struct {
	handler   BidPLacedEventHandler
	preserver PreserveBidEvent
	sender    Sender
}

func CreateBidPlacedReqCommand(conf config.CommandService) BidPlaceRequestedCommand {
	handler := command.PlaceBidHandler{
		StateRepo: adapter.CreateMariaDbStateRepository(),
		BidRepo:   adapter.CreateMariaDbBidRepository(),
	}
	preserver := command.SaveBidEventHandler{
		Repo: adapter.CreateMariaDbBidRepository(),
	}
	sender := port.CreatePublisher(conf.RedisConf.WriteUrl, port.FakeLogger{}, port.BidChannel)

	return CreateBidPlacedReqWithInterfaces(handler, preserver, sender)
}

func CreateBidPlacedReqWithInterfaces(
	handler BidPLacedEventHandler,
	preserver PreserveBidEvent,
	sender Sender,
) BidPlaceRequestedCommand {
	return BidPlaceRequestedCommand{
		handler:   handler,
		preserver: preserver,
		sender:    sender,
	}
}

func (b BidPlaceRequestedCommand) Execute(event domain.Event) {
	var bidPlacedMessage domain.BidPlaced
	notifyEvent := domain.NotifyEvent{
		Event:         event.Event,
		CorrelationId: event.CorrelationId,
	}

	if err := json.Unmarshal(event.Payload, &bidPlacedMessage); err != nil {
		notifyEvent.Error = fmt.Sprintf("can not unarshal event: %v", err)
		b.sender.NotifyUserFail(notifyEvent)
	}

	if err := b.preserver.Handle(event.Event, bidPlacedMessage); err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with saving data: %v", err)
		b.sender.NotifyUserFail(notifyEvent)
	}

	if err := b.handler.Handle(context.Background(), bidPlacedMessage); err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with bid placing: %v", err)
		b.sender.NotifyUserFail(notifyEvent)
	}

	b.sender.NotifyUserSuccess(notifyEvent)
	b.sender.PublishData(event)
}
