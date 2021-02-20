package main

import (
	"context"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/pubsub"
	"log"
)

var ctx = context.Background()

type Bid struct {
	AuctionId int    `json:"AuctionBid"`
	UserId    int    `json:"UserId"`
	Value     int    `json:"Value"`
	Action    string `json:"Action"`
}

type CreateAuction struct {
	Name      string
	DueDate   int
	StartDate int
	Timestamp int
}

type Event struct {
	Event   string
	Payload []byte
}

func main() {
	publish()
}

func publish() {
	p, err := pubsub.CreatePublisher("redis://localhost:6364")

	if err != nil {
		log.Fatal(fmt.Sprintf("can not create publisher: %v", err))
	}

	/*
		auction := CreateAuction{
			DueDate:   int(time.Now().AddDate(0, 0, 2).Unix()),
			StartDate: int(time.Now().AddDate(0, 0, 4).Unix()),
			Timestamp: int(time.Now().Unix()),
			Name:      "Test2",
		}
	*/
	/*
		user := domain.CreateUserRequested{
			Name:     "Johanna2",
			Password: "Secret",
		}

		userIvan := domain.CreateUserRequested{
			Name:     "Ivan2",
			Password: "Top Secret",
		}

		userME := domain.CreateUserRequested{
			Name:     "Gyorgy2",
			Password: "Top Secret",
		}
	*/
	/*
		userME := domain.DeleteUserRequest{
			Name: "Gyorgy2",
			Id:   6,
		}
		/*
			winner := domain.WinnerAnnounced{
				WinnerId:  2,
				AuctionId: "104b573c-cc10-418c-b9e8-64291ea720be",
			}
	*/
	/*
		bidPLaced := domain.BidPlaced{
			Promoted:  false,
			Amount:    150,
			UserId:    1,
			AuctionId: "cff3e43f-d251-49fc-a779-b57f1d87a8fe",
		}

	*/

	bidDeleted := domain.BidDeleted{
		BidId:     30,
		Amount:    100,
		UserId:    2,
		AuctionId: "cff3e43f-d251-49fc-a779-b57f1d87a8fe",
	}

	send(bidDeleted, domain.BidDeleteRequested, "Bid", p)
}

func send(message interface{}, eventName, channel string, p *pubsub.Publisher) {
	err := p.SendEvent(message, channel, eventName)

	if err != nil {
		fmt.Printf("Can not send event: %v", err)
	}
}
