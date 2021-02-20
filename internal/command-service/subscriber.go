package command_service

import (
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/adapter"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/command-service/event-reaction"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"github.com/lantosgyuri/auction-portal/internal/pkg/pubsub"
	"sync"
)

func StartSubscriber(url string, parentWg *sync.WaitGroup) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	eventChannel := make(chan domain.Event)
	eventSubscriber, err := pubsub.CreateSubscriber(url)

	if err != nil {
		fmt.Printf("error happened during create subscriber: %v", err)
		wg.Done()
	}

	eventSubscriber.AddChannel("Auction")
	eventSubscriber.AddChannel("Bid")
	eventSubscriber.AddChannel("User")

	eventSubscriber.Get(eventChannel)

	go consumeMessages(eventChannel)

	wg.Wait()
	parentWg.Done()
}

func consumeMessages(eventChan chan domain.Event) {
	application := app.Application{
		Commands: app.Commands{
			CreateAuction:    command.CreateAuctionHandler{Repo: adapter.MariaDbAuctionRepository{Db: connection.SotDb}},
			SaveAuctionEvent: command.SaveAuctionEventHandler{Repo: adapter.MariaDbAuctionRepository{Db: connection.SotDb}},
			CreateUser:       command.CreateUserHandler{Repo: adapter.MariaDbUserRepository{Db: connection.SotDb}},
			SaveUserEvent:    command.SaveUserEventHandler{Repo: adapter.MariaDbUserRepository{Db: connection.SotDb}},
			DeleteUser:       command.DeleteUserHandler{Repo: adapter.MariaDbUserRepository{Db: connection.SotDb}},
			AnnounceWinner:   command.AnnounceWinnerHandler{Repo: adapter.MariaDbStateRepository{Db: connection.SotDb}},
			SaveBidEvent:     command.SaveBidEventHandler{Repo: adapter.MariaDbBidRepository{Db: connection.SotDb}},
			PlaceBid: command.PlaceBidHandler{
				BidRepo:   adapter.MariaDbBidRepository{Db: connection.SotDb},
				StateRepo: adapter.MariaDbStateRepository{Db: connection.SotDb}},
			DeleteBid: command.DeleteBidHandler{
				BidRepo:   adapter.MariaDbBidRepository{Db: connection.SotDb},
				StateRepo: adapter.MariaDbStateRepository{Db: connection.SotDb}},
		},
	}

	for event := range eventChan {
		reaction, found := event_reaction.Commands[event.Event]
		if !found {
			fmt.Printf("no event reaction for this event: %v", event.Event)
			continue
		}
		if err := reaction.Execute(application, event); err != nil {
			fmt.Printf("error happened during event reaction: %v", err)
		}
	}
}
