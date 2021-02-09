package command_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"sync"
)

// Events can happen:
// 	User add bid to auction
// 	User delete bid from auction
//	Auction created
//	User created
//	User deleted

// Channels:
//	Auction: Auction created, bid placed, bid deleted
//	User: User created, user deleted

type InMemoryDb struct {
}

func (i InMemoryDb) SaveAuctionEvent(event domain.NormalizedAuctionEvent) error {
	fmt.Printf("saving auction event %v \n", event)
	return nil
}

func (i InMemoryDb) CreateNewAuction(auction domain.CreateAuction) error {
	fmt.Printf("saving auction event %v \n", auction)
	return nil
}

func StartSubscriber(url string) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	eventChan := make(chan domain.Event)

	application := app.Application{
		Commands: app.Commands{
			CreateAuction:    command.CreateAuctionHandler{Repo: InMemoryDb{}},
			SaveAuctionEvent: command.SaveAuctionEventHandler{Repo: InMemoryDb{}},
		},
		Queries: app.Queries{},
	}

	go subscribeToChan(url, "Auction", &wg, eventChan)
	go consumeMessages(application, eventChan)
	wg.Wait()
}

func subscribeToChan(url string, channelName string, wg *sync.WaitGroup, eventChan chan domain.Event) {
	redisConn, err := connection.SetUpRedis(url)

	if err != nil {
		fmt.Printf("can not create redis connection: %v", err)
		wg.Done()
	}

	pubs := redisConn.Subscribe(context.Background(), channelName)
	ch := pubs.Channel()
	for msg := range ch {
		var event domain.Event
		if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
			fmt.Printf("Error happened with unmarshalling: %v", err)
		}
		eventChan <- event
	}
}

func consumeMessages(application app.Application, eventChan chan domain.Event) {
	for event := range eventChan {
		switch event.Event {
		case "AUCTION_CREATED":
			var auction domain.CreateAuction
			if err := json.Unmarshal(event.Payload, &auction); err != nil {
				fmt.Printf("Error happened with unmarshalling user: %v", err)
			}
			err := application.Commands.CreateAuction.Handle(command.CreateAuction{Auction: auction})
			if err != nil {
				fmt.Printf("Error happened with creating auction: %v", err)
			}
			err = application.Commands.SaveAuctionEvent.Handle(event)
			fmt.Printf("Event is: %v", auction)

		default:
			fmt.Printf("no event like this: %v", event.Event)
		}
	}
}
