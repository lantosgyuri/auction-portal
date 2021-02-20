package event_reaction

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

var Commands = make(map[string]Command)

type Command interface {
	Execute(application app.Application, event domain.Event) error
}
