package command_service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/shared"
)

// Events can happen:
// 	User add bid to auction
// 	User delete bid from auction
//	Auction created
//	User created
//	User deleted

// Channels:
//	Auction: Auction created, bid placed, bid deleted
//	User: User created, user deleted

// Subscribe to channels
// the same event struct is coming from the queue
// It deserialize the data and check which command should be fired
// Passes the deserialized struct to the command
func StartSubscriber(url string) {
	done := make(chan bool)

	go subscribeToChan(url, "Auction")
	<-done
}

func subscribeToChan(url string, channel string) chan bool {
	done := make(chan bool)
	redisConn := shared.SetUpRedis(url)
	pubs := redisConn.Subscribe(context.Background(), channel)
	ch := pubs.Channel()

	for msg := range ch {
		var event domain.Event
		if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
			fmt.Printf("Error happened with unmarshalling: %v", err)
		}

		switch event.Event {
		case "AUCTION_CREATED":
			var user domain.CreateAuction
			if err := json.Unmarshal(event.Payload, &user); err != nil {
				fmt.Printf("Error happened with unmarshalling user: %v", err)
			}
			// CALL THE COMMAND
			fmt.Printf("Event is: %v", user)

		default:
			fmt.Printf("no event like this: %v", event.Event)
		}
	}

	return done
}
