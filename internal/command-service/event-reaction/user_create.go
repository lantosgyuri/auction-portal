package event_reaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

func init() {
	Commands[domain.UserCreateRequested] = CreateUserCommand{}
}

type CreateUserCommand struct{}

func (c CreateUserCommand) Execute(application app.Application, event domain.Event) error {
	var userCreateRequest domain.CreateUserRequested

	if err := json.Unmarshal(event.Payload, &userCreateRequest); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling user create: %v", err))
	}

	err := application.Commands.CreateUser.Handle(userCreateRequest)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened with creating user: %v", err))
	}

	err = application.Commands.SaveUserEvent.Handle(event.Event, userCreateRequest)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened during saving the user event: %v", err))
	}

	return nil
}
