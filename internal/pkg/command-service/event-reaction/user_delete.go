package event_reaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

func init() {
	Commands[domain.UserDeleteRequested] = UserDeleteCommand{}
}

type UserDeleteCommand struct{}

func (u UserDeleteCommand) Execute(application app.Application, event domain.Event) error {
	var userDeleteRequested domain.DeleteUserRequest
	if err := json.Unmarshal(event.Payload, &userDeleteRequested); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling delete user request: %v", err))
	}

	err := application.Commands.DeleteUser.Handle(userDeleteRequested)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened with deleting user: %v", err))
	}

	err = application.Commands.SaveUserEvent.Handle(event.Event, userDeleteRequested)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened during saving the user event: %v", err))
	}
	return nil
}
