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

type UserDeleteEventHandler interface {
	Handle(ctx context.Context, event domain.DeleteUserRequest) error
}

type UserDeleteCommand struct {
	handler   UserDeleteEventHandler
	preserver PreserveUserEvent
	sender    Sender
}

func CreateUserDeleteCommand(conf config.CommandService) UserDeleteCommand {
	handler := command.DeleteUserHandler{
		Repo: adapter.CreateMariaDbUserRepository(),
	}
	preserver := command.SaveUserEventHandler{
		Repo: adapter.CreateMariaDbUserRepository(),
	}

	sender := port.CreatePublisher(conf.RedisConf.WriteUrl, port.FakeLogger{}, port.UserChannel)

	return CreateUserDeleteWithInterfaces(handler, preserver, sender)
}

func CreateUserDeleteWithInterfaces(
	handler UserDeleteEventHandler,
	preserver PreserveUserEvent,
	sender Sender,
) UserDeleteCommand {
	return UserDeleteCommand{
		handler:   handler,
		preserver: preserver,
		sender:    sender,
	}
}

func (u UserDeleteCommand) Execute(event domain.Event) {
	var userDeleteRequested domain.DeleteUserRequest
	notifyEvent := domain.NotifyEvent{
		Event:         event.Event,
		CorrelationId: event.CorrelationId,
	}
	if err := json.Unmarshal(event.Payload, &userDeleteRequested); err != nil {
		notifyEvent.Error = fmt.Sprintf("can not unarshal event: %v", err)
		u.sender.NotifyUserFail(notifyEvent)
	}

	err := u.handler.Handle(context.Background(), userDeleteRequested)
	if err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with user deleting: %v", err)
		u.sender.NotifyUserFail(notifyEvent)
	}

	err = u.preserver.Handle(event.Event, userDeleteRequested)
	if err != nil {
		notifyEvent.Error = fmt.Sprintf("error happened with saving data: %v", err)
		u.sender.NotifyUserFail(notifyEvent)
	}

	u.sender.NotifyUserSuccess(notifyEvent)
	u.sender.PublishData(event)
}
