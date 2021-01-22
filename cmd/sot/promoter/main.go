package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Bid struct {
	AuctionId int `json:"AuctionBid"`
	UserId    int `json:"UserId"`
	Value     int `json:"Value"`
}

type User struct {
	Id       int   `json:"Id"`
	Auctions []int `json:"Auctions"`
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
	var i int

	redisConn := CreateRedisClient()

	for i < 10 {
		newBid := Bid{
			AuctionId: i,
			UserId:    i + 10,
			Value:     i + 20,
		}
		fmt.Println(newBid)
		toSend, err := json.Marshal(newBid)
		if err != nil {
			fmt.Print(err)
		}
		redisConn.Publish(ctx, "bid", toSend)
		i++
	}

	for i < 20 {
		user := User{
			Id: i,
			Auctions: []int{
				i + 10,
			},
		}
		fmt.Println(user)
		toSend, err := json.Marshal(user)
		if err != nil {
			fmt.Print(err)
		}
		redisConn.Publish(ctx, "user", toSend)
		i++
	}
}
