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

	auctionEventChan := make(chan domain.Event)
	bidEventChan := make(chan domain.Event)
	userEventChan := make(chan domain.Event)

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

	go handleChannel(url, "Auction", &wg, application, auctionEventChan)
	go handleChannel(url, "Bid", &wg, application, bidEventChan)
	go handleChannel(url, "User", &wg, application, userEventChan)

	wg.Wait()
	parentWg.Done()
}

func handleChannel(url string, channelName string, wg *sync.WaitGroup, application app.Application, eventChan chan domain.Event) {
	go subscribeToChan(url, channelName, wg, eventChan)
	go consumeMessages(application, eventChan)
}

func subscribeToChan(url string, channelName string, wg *sync.WaitGroup, eventChan chan domain.Event) {
	subscriber, err := pubsub.CreateSubscriber(url)

	if err != nil {
		fmt.Printf("error happened during create subscriber: %v", err)
		wg.Done()
	}

	subscriber.GetEvents(channelName, eventChan)
}

func consumeMessages(application app.Application, eventChan chan domain.Event) {
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
