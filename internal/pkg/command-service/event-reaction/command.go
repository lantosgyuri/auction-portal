package event_reaction

import (
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

var Commands = make(map[string]Command)

var EventMessageNames = EventMessageNamesConst{
	AuctionRequested: "Auction_Creation_Requested",
}

type Command interface {
	Execute(application app.Application, event domain.Event) error
}

type EventMessageNamesConst struct {
	AuctionRequested string
}
