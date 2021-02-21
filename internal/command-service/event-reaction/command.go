package event_reaction

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

type Command interface {
	Execute(event domain.Event) error
}

func CreateCommands() map[string]Command {
	commands := make(map[string]Command)
	commands[domain.AuctionRequested] = CreateAuctionRequestedCommand()
	commands[domain.BidDeleteRequested] = CreateBidDeletedCommand()
	commands[domain.BidPlaceRequested] = CreateBidPlacedReqCommand()
	commands[domain.UserCreateRequested] = MakeCreateUserCommand()
	commands[domain.UserDeleteRequested] = CreateUserDeleteCommand()
	commands[domain.AuctionWinnerAnnounced] = CreateWinnerAnnouncedCommand()
	return commands
}
