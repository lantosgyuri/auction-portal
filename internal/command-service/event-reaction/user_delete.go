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

type UserDeleteEventHandler interface {
	Handle(ctx context.Context, event domain.DeleteUserRequest) error
}

type UserDeleteCommand struct {
	handler   UserDeleteEventHandler
	preserver PreserveUserEvent
	publisher EventPublisher
}

func CreateUserDeleteCommand() UserDeleteCommand {
	handler := command.DeleteUserHandler{
		Repo: adapter.CreateMariaDbUserRepository(),
	}
	preserver := command.SaveUserEventHandler{
		Repo: adapter.CreateMariaDbUserRepository(),
	}

	return CreateUserDeleteWithInterfaces(handler, preserver)
}

func CreateUserDeleteWithInterfaces(handler UserDeleteEventHandler, preserver PreserveUserEvent) UserDeleteCommand {
	return UserDeleteCommand{
		handler:   handler,
		preserver: preserver,
	}
}

func (u UserDeleteCommand) Execute(event domain.Event) error {
	var userDeleteRequested domain.DeleteUserRequest
	if err := json.Unmarshal(event.Payload, &userDeleteRequested); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling delete user request: %v", err))
	}

	err := u.handler.Handle(context.Background(), userDeleteRequested)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened with deleting user: %v", err))
	}

	err = u.preserver.Handle(event.Event, userDeleteRequested)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened during saving the user event: %v", err))
	}
	err = u.publisher.Publish(event)
	if err != nil {
		return errors.New(fmt.Sprintf("Can not publish event: %v", err))
	}
	return nil
}
