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

type BidDeletedEventHandler interface {
	Handle(ctx context.Context, event domain.BidDeleted) error
}

type BidDeleteRequestedCommand struct {
	handler   BidDeletedEventHandler
	preserver PreserveBidEvent
	sender    Sender
}

func CreateBidDeletedCommand(conf config.CommandService) BidDeleteRequestedCommand {
	handler := command.DeleteBidHandler{
		BidRepo:   adapter.CreateMariaDbBidRepository(),
		StateRepo: adapter.CreateMariaDbStateRepository(),
	}
	preserver := command.SaveBidEventHandler{
		Repo: adapter.CreateMariaDbBidRepository(),
	}
	sender := port.CreatePublisher(conf.RedisConf.QueryUrl, port.FakeLogger{}, port.BidChannel)

	return CreateBidDeletedWithInterfaces(handler, preserver, sender)
}

func CreateBidDeletedWithInterfaces(
	handler BidDeletedEventHandler,
	preserver PreserveBidEvent,
	sender Sender,
) BidDeleteRequestedCommand {
	return BidDeleteRequestedCommand{
		handler:   handler,
		preserver: preserver,
		sender:    sender,
	}
}

func (b BidDeleteRequestedCommand) Execute(event domain.Event) {
	var bidDeleteMessage domain.BidDeleted
	notifyEvent := domain.NotifyEvent{
		Event:         event.Event,
		CorrelationId: event.CorrelationId,
	}
	if err := json.Unmarshal(event.Payload, &bidDeleteMessage); err != nil {
		notifyEvent.Error = fmt.Sprintf("can not unarshal event: %v", err)
		b.sender.NotifyUserFail(notifyEvent)
		return
	}

	if err := b.preserver.Handle(event.Event, bidDeleteMessage); err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with saving data: %v", err)
		b.sender.NotifyUserFail(notifyEvent)
		return
	}

	if err := b.handler.Handle(context.Background(), bidDeleteMessage); err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with deleting bid: %v", err)
		b.sender.NotifyUserFail(notifyEvent)
		return
	}

	notifyEvent.Success = true
	b.sender.NotifyUserSuccess(notifyEvent)
	b.sender.PublishData(event)
}
