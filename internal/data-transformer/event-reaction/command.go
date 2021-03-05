package event_reaction

import (
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/data-transformer/app/command"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
)

type Preserver interface {
	Execute(event domain.Event) error
}

type Command struct {
	Preserver Preserver
}

func (c *Command) Do(event domain.Event) {
	err := c.Preserver.Execute(event)

	if err != nil {
		fmt.Printf("can not save data: %v", err)
	}
}

type Repo struct{}

func (r Repo) SaveAuction(auction domain.CreateAuctionRequested) error {
	fmt.Printf("SAVING AUCTION %v", auction)
	return nil
}

func (r Repo) SaveWinner(winner domain.WinnerAnnounced) error {
	fmt.Printf("SAVING WINNER %v", winner)
	return nil
}

func (r Repo) SaveBid(bid domain.BidPlaced) error {
	fmt.Printf("SAVING BID %v", bid)
	return nil
}

func (r Repo) DeleteBid(bid domain.BidDeleted) error {
	fmt.Printf("DELETING BID %v", bid)
	return nil
}

func (r Repo) SaveUser(user domain.CreateUserRequested) error {
	fmt.Printf("SAVING USER %v", user)
	return nil
}

func (r Repo) DeleteUser(user domain.DeleteUserRequest) error {
	fmt.Printf("DELETING USER %v", user)
	return nil
}

func CreateCommands(conf config.DataTransformer) map[string]Command {
	commands := make(map[string]Command)

	commands[domain.AuctionRequested] = Command{
		Preserver: &command.AuctionPreserver{
			Repo: Repo{},
		},
	}

	commands[domain.AuctionWinnerAnnounced] = Command{
		Preserver: &command.SaveWinner{
			Repo: Repo{},
		},
	}

	commands[domain.BidPlaceRequested] = Command{
		Preserver: &command.SaveBid{
			Repo: Repo{},
		},
	}

	commands[domain.BidDeleteRequested] = Command{
		Preserver: &command.DeleteBid{
			Repo: Repo{},
		},
	}

	commands[domain.UserCreateRequested] = Command{
		Preserver: &command.SaveUser{
			Repo: Repo{},
		},
	}

	commands[domain.UserDeleteRequested] = Command{
		Preserver: &command.DeleteUser{
			Repo: Repo{},
		},
	}

	return commands
}
