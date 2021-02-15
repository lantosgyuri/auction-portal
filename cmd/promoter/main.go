package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
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

func CreateRedisClient() *redis.Client {
	opt, err := redis.ParseURL("redis://localhost:6364")
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}

func publish() {
	redisConn := CreateRedisClient()

	user := CreateAuction{
		DueDate:   int(time.Now().AddDate(0, 0, 7).Unix()),
		StartDate: int(time.Now().AddDate(0, 0, 2).Unix()),
		Timestamp: int(time.Now().Unix()),
		Name:      "JOZSIKA",
	}

	userBytes, _ := json.Marshal(user)

	event := Event{
		Event:   domain.AuctionRequested,
		Payload: userBytes,
	}

	messageBytes, _ := json.Marshal(event)

	redisConn.Publish(ctx, "Auction", messageBytes)
	/*
		winner := domain.WinnerAnnounced{
			Timestamp: int(time.Now().Unix()),
			WinnerId:  19880,
			AuctionId: "test",
		}

		winnerBytes, _ := json.Marshal(winner)

		eventWinner := Event{
			Event:   domain.AuctionWinnerAnnounced,
			Payload: winnerBytes,
		}

		winnerMessageBytes, _ := json.Marshal(eventWinner)

		redisConn.Publish(ctx, "Auction", winnerMessageBytes)

		bidPLaced := domain.BidPlaced{
			Promoted:  false,
			Amount:    20,
			UserId:    41,
			AuctionId: "test",
		}

		bidPLacedbytes, _ := json.Marshal(bidPLaced)

		bidPlacedEvent := Event{
			Event:   domain.BidPlaceRequested,
			Payload: bidPLacedbytes,
		}

		bidPLacedEvetnBytes, _ := json.Marshal(bidPlacedEvent)

		redisConn.Publish(ctx, "Bid", bidPLacedEvetnBytes)

		bidDeleted := domain.BidDeleted{
			BidId:     40,
			Amount:    20,
			UserId:    41,
			AuctionId: "test",
		}

		bidDeletedbytes, _ := json.Marshal(bidDeleted)

		bidDeletedEcent := Event{
			Event:   domain.BidDeleteRequested,
			Payload: bidDeletedbytes,
		}

		bidDeletedEvetnBytes, _ := json.Marshal(bidDeletedEcent)

		redisConn.Publish(ctx, "Bid", bidDeletedEvetnBytes)
	*/
}
