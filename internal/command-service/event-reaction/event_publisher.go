package event_reaction

import "github.com/lantosgyuri/auction-portal/internal/command-service/domain"

type EventPublisher interface {
	Publish(event domain.Event) error
}
