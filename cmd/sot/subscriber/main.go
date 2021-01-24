package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/shared"
	"gopkg.in/yaml.v3"
	"log"
)

var ctx = context.Background()

type Bid struct {
	AuctionId int    `json:"AuctionBid"`
	UserId    int    `json:"UserId"`
	Value     int    `json:"Value"`
	Action    string `json:"Action"`
}

type User struct {
	Id       int   `json:"Id"`
	Auctions []int `json:"Auctions"`
}

type Config struct {
	RedisConf redisConf `yaml:"redis"`
}

type redisConf struct {
	Url  string   `yaml:"url"`
	Subs []string `yaml:"subscriptions"`
}

func main() {
	done := make(chan bool)
	var conf Config

	confBytes, err := shared.ReadFile("config.yaml")

	if err != nil {
		log.Fatal("Can not read config file")
	}

	err = yaml.Unmarshal(confBytes, &conf)

	if err != nil {
		log.Fatal("Can not read config file")
	}

	fmt.Print(conf.RedisConf)
	go subscribeToChan(conf.RedisConf.Url, conf.RedisConf.Subs[0])
	go subscribeToChan2(conf.RedisConf.Url, conf.RedisConf.Subs[1])
	<-done
}

func subscribeToChan(url string, channel string) chan bool {
	done := make(chan bool)
	redisConn := shared.SetUpRedis(url)
	pubs := redisConn.Subscribe(ctx, channel)
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

func subscribeToChan2(url string, channel string) chan bool {
	done := make(chan bool)
	redisConn := shared.SetUpRedis(url)
	pubs := redisConn.Subscribe(ctx, channel)
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
