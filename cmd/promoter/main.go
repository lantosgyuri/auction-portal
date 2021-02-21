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
			DueDate:   int(time.Now().AddDate(0, 0, 8).Unix()),
			StartDate: int(time.Now().AddDate(0, 0, 1).Unix()),
			Timestamp: int(time.Now().Unix()),
			Name:      "Test3",
		}
	*/
	/*
		user := domain.CreateUserRequested{
			Name:     "Mia",
			Password: "Secret",
		}

	*/
	/*
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
					Name: "Mia",
					Id:   7,
				}


			winner := domain.WinnerAnnounced{
				WinnerId:  2,
				AuctionId: "0bd37b85-f5d1-4418-a796-7eaf29980005",
			}


		bidPLaced := domain.BidPlaced{
			Promoted:  false,
			Amount:    250,
			UserId:    3,
			AuctionId: "0bd37b85-f5d1-4418-a796-7eaf29980005",
		}
	*/

	bidDeleted := domain.BidDeleted{
		BidId:     38,
		Amount:    250,
		UserId:    3,
		AuctionId: "0bd37b85-f5d1-4418-a796-7eaf29980005",
	}

	send(bidDeleted, domain.BidDeleteRequested, "Bid", p)

}

func send(message interface{}, eventName, channel string, p *pubsub.Publisher) {
	err := p.SendEvent(message, channel, eventName)

	if err != nil {
		fmt.Printf("Can not send event: %v", err)
	}
}
