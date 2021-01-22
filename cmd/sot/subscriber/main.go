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
	done := make(chan bool)
	go subscribeToChan()
	go subscribeToChan2()
	<-done
}

func CreateRedisClient() *redis.Client {
	opt, err := redis.ParseURL("redis://localhost:6364")
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opt)
}

func subscribeToChan() chan bool {
	done := make(chan bool)
	redisConn := CreateRedisClient()
	pubs := redisConn.Subscribe(ctx, "Test")
	ch := pubs.Channel()

	for msg := range ch {
		fmt.Print("getting message \n")
		var bid Bid
		if err := json.Unmarshal([]byte(msg.Payload), &bid); err != nil {
			fmt.Printf("Error happened with unmarshalling: %v", err)
		}
		fmt.Printf("Message is : %v", bid)
	}

	return done
}

func subscribeToChan2() chan bool {
	done := make(chan bool)
	redisConn := CreateRedisClient()
	pubs := redisConn.Subscribe(ctx, "Test2")
	ch := pubs.Channel()

	for msg := range ch {
		fmt.Print("getting message \n")
		var user User
		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			fmt.Printf("Error happened with unmarshalling: %v", err)
		}
		fmt.Printf("Message is : %v", user)
	}

	return done
}
