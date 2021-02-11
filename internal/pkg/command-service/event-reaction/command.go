package event_reaction

import (
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

var Commands = make(map[string]Command)

type Command interface {
	Execute(application app.Application, event domain.Event) error
}
