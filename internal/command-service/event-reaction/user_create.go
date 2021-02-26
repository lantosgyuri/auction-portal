package event_reaction

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/adapter"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
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
	publisher EventPublisher
}

func MakeCreateUserCommand() CreateUserCommand {
	handler := command.CreateUserHandler{
		Repo: adapter.CreateMariaDbUserRepository(),
	}
	preserver := command.SaveUserEventHandler{
		Repo: adapter.CreateMariaDbUserRepository(),
	}

	return MakeCreateUserWithInterfaces(handler, preserver)
}

func MakeCreateUserWithInterfaces(handler CreateUserEventHandler, preserver PreserveUserEvent) CreateUserCommand {
	return CreateUserCommand{
		handler:   handler,
		preserver: preserver,
	}
}

func (c CreateUserCommand) Execute(event domain.Event) error {
	var userCreateRequest domain.CreateUserRequested

	if err := json.Unmarshal(event.Payload, &userCreateRequest); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling user create: %v", err))
	}

	err := c.handler.Handle(context.Background(), userCreateRequest)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened with creating user: %v", err))
	}

	err = c.preserver.Handle(event.Event, userCreateRequest)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened during saving the user event: %v", err))
	}
	err = c.publisher.Publish(event)
	if err != nil {
		return errors.New(fmt.Sprintf("Can not publish event: %v", err))
	}
	return nil
}
