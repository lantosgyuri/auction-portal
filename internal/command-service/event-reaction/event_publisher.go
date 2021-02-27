package event_reaction

import "github.com/lantosgyuri/auction-portal/internal/command-service/domain"

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
