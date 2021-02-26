package event_reaction

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
)

type Command interface {
	Execute(event domain.Event) error
}

func CreateCommands(conf config.CommandService) map[string]Command {
	commands := make(map[string]Command)
	commands[domain.AuctionRequested] = CreateAuctionRequestedCommand(conf)
	commands[domain.BidDeleteRequested] = CreateBidDeletedCommand(conf)
	commands[domain.BidPlaceRequested] = CreateBidPlacedReqCommand(conf)
	commands[domain.UserCreateRequested] = MakeCreateUserCommand(conf)
	commands[domain.UserDeleteRequested] = CreateUserDeleteCommand(conf)
	commands[domain.AuctionWinnerAnnounced] = CreateWinnerAnnouncedCommand(conf)
	return commands
}
