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

	winner := domain.AuctionWinnerMessage{
		Timestamp: int(time.Now().Unix()),
		WinnerId:  1,
		AuctionId: 1,
	}

	winnerBytes, _ := json.Marshal(winner)

	eventWinner := Event{
		Event:   domain.AuctionWinnerAnnounced,
		Payload: winnerBytes,
	}

	winnerMessageBytes, _ := json.Marshal(eventWinner)

	redisConn.Publish(ctx, "Auction", winnerMessageBytes)
}
