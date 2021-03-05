package pubsub

import (
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

func SubscribeToMainEvents(subscriber *Subscriber, eventChannel chan domain.Event) {
	messageChannel := make(chan []byte)

	subscriber.AddChannel("Auction")
	subscriber.AddChannel("Bid")
	subscriber.AddChannel("User")

	subscriber.Get(messageChannel)
	go convertMainEventMessage(messageChannel, eventChannel)
}

func convertMainEventMessage(messageChannel chan []byte, eventChannel chan domain.Event) {
	for message := range messageChannel {
		var event domain.Event
		if err := json.Unmarshal(message, &event); err != nil {
			fmt.Printf("error happened during unmarshal event: %v", err)
		}
		eventChannel <- event
	}
}
