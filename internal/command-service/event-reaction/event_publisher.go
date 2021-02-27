package event_reaction

import "github.com/lantosgyuri/auction-portal/internal/command-service/domain"

type UserNotifier interface {
	NotifyUserSuccess(correlationID int, event string)
	NotifyUserFail(correlationID int, event string, err error)
}

type DataPublisher interface {
	PublishData(event domain.Event)
}

type Sender interface {
	UserNotifier
	DataPublisher
}
