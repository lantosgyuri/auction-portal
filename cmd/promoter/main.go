package main

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
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
		DueDate:   6,
		StartDate: 3,
		Timestamp: 1,
		Name:      "JOZSIKA",
	}

	userBytes, _ := json.Marshal(user)

	event := Event{
		Event:   "AUCTION_CREATED",
		Payload: userBytes,
	}

	messageBytes, _ := json.Marshal(event)

	redisConn.Publish(ctx, "Auction", messageBytes)
}
