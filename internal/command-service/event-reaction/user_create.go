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

type CreateUserEventHandler interface {
	Handle(ctx context.Context, event domain.CreateUserRequested) error
}

type PreserveUserEvent interface {
	Handle(eventName string, event domain.UserEvent) error
}

type CreateUserCommand struct {
	handler   CreateUserEventHandler
	preserver PreserveUserEvent
	sender    Sender
}

func MakeCreateUserCommand(conf config.CommandService) CreateUserCommand {
	handler := command.CreateUserHandler{
		Repo: adapter.CreateMariaDbUserRepository(),
	}
	preserver := command.SaveUserEventHandler{
		Repo: adapter.CreateMariaDbUserRepository(),
	}
	sender := port.CreatePublisher(conf.RedisConf.WriteUrl, port.FakeLogger{}, port.UserChannel)

	return MakeCreateUserWithInterfaces(handler, preserver, sender)
}

func MakeCreateUserWithInterfaces(
	handler CreateUserEventHandler,
	preserver PreserveUserEvent,
	sender Sender,
) CreateUserCommand {
	return CreateUserCommand{
		handler:   handler,
		preserver: preserver,
		sender:    sender,
	}
}

func (c CreateUserCommand) Execute(event domain.Event) {
	var userCreateRequest domain.CreateUserRequested
	notifyEvent := domain.NotifyEvent{
		Event:         event.Event,
		CorrelationId: event.CorrelationId,
	}
	if err := json.Unmarshal(event.Payload, &userCreateRequest); err != nil {
		notifyEvent.Error = fmt.Sprintf("can not unarshal event: %v", err)
		c.sender.NotifyUserFail(notifyEvent)
		return
	}

	err := c.preserver.Handle(event.Event, userCreateRequest)
	if err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with saving data: %v", err)
		c.sender.NotifyUserFail(notifyEvent)
		return
	}

	err = c.handler.Handle(context.Background(), userCreateRequest)
	if err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with user creating: %v", err)
		c.sender.NotifyUserFail(notifyEvent)
		return
	}

	notifyEvent.Success = true
	c.sender.NotifyUserSuccess(notifyEvent)
	c.sender.PublishData(event)
}
