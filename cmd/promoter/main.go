package main

import (
	"context"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/pubsub"
	"log"
	"time"
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

	auction := CreateAuction{
		DueDate:   int(time.Now().AddDate(0, 0, 8).Unix()),
		StartDate: int(time.Now().AddDate(0, 0, 1).Unix()),
		Timestamp: int(time.Now().Unix()),
		Name:      "Test6",
	}

	send(auction, domain.AuctionRequested, "Auction", p)
	/*
				user := domain.CreateUserRequested{
					Name:     "Mia",
					Password: "Secret",
				}


				/*
					userIvan := domain.CreateUserRequested{
						Name:     "Ivan2",
						Password: "Top Secret",
					}

					userME := domain.CreateUserRequested{
						Name:     "Gyorgy2",
						Password: "Top Secret",
					}


			userME := domain.DeleteUserRequest{
				Name: "Mia",
				Id:   8,
			}



		winner := domain.WinnerAnnounced{
			WinnerId:  2,
			AuctionId: "00216df7-086b-4d47-b350-ca4c37ca47ab",
		}
	*/
	bidPLaced := domain.BidPlaced{
		Promoted:  false,
		Amount:    5901,
		UserId:    3,
		AuctionId: "7dc73f8b-f69e-49a5-81e2-eb74814d4251",
	}

	fmt.Println(bidPLaced)
	//send(bidPLaced, domain.BidPlaceRequested, "Bid", p)

	bidDeleted := domain.BidDeleted{
		BidId:     54,
		Amount:    5901,
		UserId:    3,
		AuctionId: "7dc73f8b-f69e-49a5-81e2-eb74814d4251",
	}

	fmt.Print(bidDeleted)

}

func send(message interface{}, eventName, channel string, p *pubsub.Publisher) {
	err := p.SendEvent(message, channel, eventName)

	if err != nil {
		fmt.Printf("Can not send event: %v", err)
	}
}
