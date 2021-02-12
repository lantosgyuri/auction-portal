package command_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/query"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/event-reaction"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"sync"
)

type InMemoryDb struct {
}

func (i InMemoryDb) SaveAuctionEvent(event domain.AuctionEventRaw) error {
	fmt.Printf("saving auction event %v \n", event)
	return nil
}

func (i InMemoryDb) CreateNewAuction(auction domain.Auction) error {
	fmt.Printf("saving auction %v \n", auction)
	return nil
}

func (i InMemoryDb) UpdateAuctionState(auction domain.Auction) error {
	fmt.Printf("updating auction %v \n", auction)
	return nil
}

func (i InMemoryDb) SaveBidEvent(event domain.BidEventRaw) error {
	fmt.Printf("saving bid event %v \n", event)
	return nil
}

func (I InMemoryDb) FindAuction(ctx context.Context, auctionID string) (domain.Auction, error) {
	return domain.Auction{
		UUID:        "t",
		Name:        "Test",
		StartDate:   1,
		DueDate:     2,
		CurrentBid:  0,
		CurrentUser: 0,
	}, nil
}

func StartSubscriber(url string, parentWg *sync.WaitGroup) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	auctionEventChan := make(chan domain.Event)
	bidEventChan := make(chan domain.Event)
	userEventChan := make(chan domain.Event)

	application := app.Application{
		Commands: app.Commands{
			CreateAuction:    command.CreateAuctionHandler{Repo: InMemoryDb{}},
			SaveAuctionEvent: command.SaveAuctionEventHandler{Repo: InMemoryDb{}},
			UpdateState:      command.UpdateStateHandler{Repo: InMemoryDb{}},
			SaveBidEvent:     command.SaveBidEventHandler{Repo: InMemoryDb{}},
		},
		Queries: app.Queries{
			GetAuctionState: query.AuctionStateHandler{
				Reader: InMemoryDb{},
			},
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
