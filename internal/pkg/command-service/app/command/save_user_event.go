package command

import "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"

type SaveUserEventHandler struct {
	Repo UserRepository
}

func (s SaveUserEventHandler) Handle(eventName string, event domain.UserEvent) error {
	rawEvent := domain.UserEventRaw{}

	switch e := event.(type) {
	case domain.CreateUserRequested:
		rawEvent.EventType = eventName
		rawEvent.Name = e.Name
		rawEvent.Password = e.Password

	case domain.DeleteUserRequest:
		rawEvent.EventType = eventName
		rawEvent.Name = e.Name
		rawEvent.UserId = e.Id
	}

	return s.Repo.SaveUserEvent(rawEvent)
}
