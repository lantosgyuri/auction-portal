package event_reaction

import (
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/data-transformer/app/command"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
)

type Preserver interface {
	Execute(event Event) error
}

type Command struct {
	Preserver Preserver
}

func (c *Command) Do(event Event) {
	err := c.Preserver.Execute(event)

	if err != nil {
		fmt.Printf("can not save data: %v", err)
	}
}

type Repo struct{}

func (r Repo) SaveAuction(auction AuctionCreatedEvent) error {
	fmt.Printf("SAVING AUCTION %v", auction)
	return nil
}

func (r Repo) UpdateAuctionBid(bid BidPlacedEvent) error {
	fmt.Printf("UPDATE AUCTION %v", bid)
	return nil
}

func (r Repo) UpdateAuctionWinner(winner WinnerAnnouncedEvent) error {
	fmt.Printf("UPDATE AUCTION %v", winner)
	return nil
}

func (r Repo) SaveWinner(winner WinnerAnnouncedEvent) error {
	fmt.Printf("SAVING WINNER %v", winner)
	return nil
}

func (r Repo) SaveUserBid(bid BidPlacedEvent) error {
	fmt.Printf("SAVING BID %v", bid)
	return nil
}

func (r Repo) DeleteUserBid(bid BidDeletedEvent) error {
	fmt.Printf("DELETING BID %v", bid)
	return nil
}

func (r Repo) SaveUser(user UserCreatedEvent) error {
	fmt.Printf("SAVING USER %v", user)
	return nil
}

func (r Repo) DeleteUser(user UserDeletedEvent) error {
	fmt.Printf("DELETING USER %v", user)
	return nil
}

func (r Repo) CreateUser(user UserCreatedEvent) error {
	fmt.Printf("CREATE BID USER %v", user)
	return nil
}

func (r Repo) DeleteUserEntries(user UserDeletedEvent) error {
	fmt.Printf("DELETE BID USER %v", user)
	return nil
}

func CreateCommands(conf config.DataTransformer) map[string]Command {
	commands := make(map[string]Command)

	commands[AuctionCreated] = Command{
		Preserver: &command.AuctionPreserver{
			AuctionRepo: Repo{},
		},
	}

	commands[WinnerAnnounced] = Command{
		Preserver: &command.SaveWinner{
			AuctionRepo: Repo{},
			WinnerRepo:  Repo{},
		},
	}

	commands[BidPlaced] = Command{
		Preserver: &command.SaveBid{
			BidRepo:     Repo{},
			AuctionRepo: Repo{},
		},
	}

	commands[BidDeleted] = Command{
		Preserver: &command.DeleteBid{
			BidRepo: Repo{},
		},
	}

	commands[UserCreated] = Command{
		Preserver: &command.SaveUser{
			UserRepo: Repo{},
			BidRepo:  Repo{},
		},
	}

	commands[UserDeleted] = Command{
		Preserver: &command.DeleteUser{
			UserRepo: Repo{},
			BidRepo:  Repo{},
		},
	}

	return commands
}
