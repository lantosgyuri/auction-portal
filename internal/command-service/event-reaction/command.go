package event_reaction

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
	"github.com/mitchellh/mapstructure"
)

type Command interface {
	Execute(event domain.Event)
}

func CreateCommands(conf config.CommandService) map[string]Command {
	commands := make(map[string]Command)
	commands[domain.AuctionRequested] = CreateAuctionRequestedCommand(conf)
	commands[domain.BidDeleteRequested] = CreateBidDeletedCommand(conf)
	commands[domain.BidPlaceRequested] = CreateBidPlacedReqCommand(conf)
	commands[domain.UserCreateRequested] = MakeCreateUserCommand(conf)
	commands[domain.UserDeleteRequested] = CreateUserDeleteCommand(conf)
	commands[domain.AuctionWinnerAnnounced] = CreateWinnerAnnouncedCommand(conf)
	return commands
}

type Handler interface {
	Handle(context context.Context, event interface{}) error
}

type Preserver interface {
	Preserve(eventName string, event interface{}) error
}

type UserNotifier interface {
	NotifyUserSuccess(notifyEvent domain.NotifyEvent)
	NotifyUserFail(notifyEvent domain.NotifyEvent)
}

type DataPublisher interface {
	PublishData(event domain.Event)
}

type Sender interface {
	UserNotifier
	DataPublisher
}

type EventReactor struct {
	handler   Handler
	preserver Preserver
	sender    Sender
}

func (e EventReactor) Execute(event domain.Event) {
	notifyEvent := domain.NotifyEvent{
		Event:         event.Event,
		CorrelationId: event.CorrelationId,
	}

	var tmp interface{}

	if err := json.Unmarshal(event.Payload, &tmp); err != nil {
		notifyEvent.Error = fmt.Sprintf("can not unarshal event: %v", err)
		e.sender.NotifyUserFail(notifyEvent)
		return
	}

	err := mapstructure.Decode(tmp, &e.message)

	err = e.preserver.Preserve(event.Event, e.message)
	if err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with saving data: %v", err)
		e.sender.NotifyUserFail(notifyEvent)
		return
	}

	err = e.handler.Handle(context.Background(), e.message)
	if err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with auction creating: %v", err)
		e.sender.NotifyUserFail(notifyEvent)
		return
	}

	e.sender.NotifyUserSuccess(notifyEvent)
	e.sender.PublishData(event)
}
