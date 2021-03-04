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

type WinnerAnnouncedEventHandler interface {
	Handle(context context.Context, event domain.WinnerAnnounced) error
}

type WinnerAnnouncedCommand struct {
	handler   WinnerAnnouncedEventHandler
	preserver AuctionEventPreserver
	sender    Sender
}

func CreateWinnerAnnouncedCommand(conf config.CommandService) WinnerAnnouncedCommand {
	handler := command.AnnounceWinnerHandler{Repo: adapter.CreateMariaDbStateRepository()}
	preserver := command.SaveAuctionEventHandler{Repo: adapter.CreateMariaDbAuctionRepository()}
	sender := port.CreatePublisher(conf.RedisConf.WriteUrl, port.FakeLogger{}, port.AuctionChannel)
	return CreateWinnerCommandWithInterfaces(handler, preserver, sender)
}

func CreateWinnerCommandWithInterfaces(
	handler WinnerAnnouncedEventHandler,
	preserver AuctionEventPreserver,
	sender Sender,
) WinnerAnnouncedCommand {
	return WinnerAnnouncedCommand{
		handler:   handler,
		preserver: preserver,
		sender:    sender,
	}
}

func (w WinnerAnnouncedCommand) Execute(event domain.Event) {
	var winnerMessage domain.WinnerAnnounced
	notifyEvent := domain.NotifyEvent{
		Event:         event.Event,
		CorrelationId: event.CorrelationId,
	}
	if err := json.Unmarshal(event.Payload, &winnerMessage); err != nil {
		notifyEvent.Error = fmt.Sprintf("can not unarshal event: %v", err)
		w.sender.NotifyUserFail(notifyEvent)
		return
	}

	if err := w.preserver.Handle(event.Event, winnerMessage); err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with saving data: %v", err)
		w.sender.NotifyUserFail(notifyEvent)
		return
	}

	if err := w.handler.Handle(context.Background(), winnerMessage); err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with winner announcing: %v", err)
		w.sender.NotifyUserFail(notifyEvent)
		return
	}

	notifyEvent.Success = true
	w.sender.NotifyUserSuccess(notifyEvent)
	w.sender.PublishData(event)
}
