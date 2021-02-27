package command_service

import (
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/command-service/event-reaction"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
	"github.com/lantosgyuri/auction-portal/internal/pkg/pubsub"
	"sync"
)

func StartSubscriber(conf config.CommandService, parentWg *sync.WaitGroup) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	messageChannel := make(chan []byte)
	eventChannel := make(chan domain.Event)
	eventSubscriber, err := pubsub.CreateSubscriber(conf.RedisConf.WriteUrl)

	if err != nil {
		fmt.Printf("error happened during create subscriber: %v", err)
		wg.Done()
	}

	eventSubscriber.AddChannel("Auction")
	eventSubscriber.AddChannel("Bid")
	eventSubscriber.AddChannel("User")

	eventSubscriber.Get(messageChannel)
	go convertMessage(messageChannel, eventChannel)
	go consumeMessages(conf, eventChannel)

	wg.Wait()
	parentWg.Done()
}

func convertMessage(messageChannel chan []byte, eventChannel chan domain.Event) {
	for message := range messageChannel {
		var event domain.Event
		if err := json.Unmarshal(message, &event); err != nil {
			fmt.Printf("error happened during unmarshal event: %v", err)
		}
		eventChannel <- event
	}
}

func consumeMessages(config config.CommandService, eventChan chan domain.Event) {
	commands := event_reaction.CreateCommands(config)
	for event := range eventChan {
		reaction, found := commands[event.Event]
		if !found {
			fmt.Printf("no event reaction for this event: %v", event.Event)
			continue
		}
		reaction.Execute(event)
	}
}
