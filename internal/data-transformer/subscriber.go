package data_transformer

import (
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	event_reaction "github.com/lantosgyuri/auction-portal/internal/data-transformer/event-reaction"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
	"github.com/lantosgyuri/auction-portal/internal/pkg/pubsub"
	"sync"
)

func StartSubscriber(conf config.DataTransformer, parentWg *sync.WaitGroup) {
	eventChannel := make(chan domain.Event)
	subscriber, err := pubsub.CreateSubscriber(conf.RedisConf.QueueUrl)

	if err != nil {
		fmt.Printf("error happened during create subscriber: %v", err)
		parentWg.Done()
	}

	go pubsub.SubscribeToMainEvents(subscriber, eventChannel)
	go consumeMessages(conf, eventChannel)
}

func consumeMessages(conf config.DataTransformer, eventChannel chan domain.Event) {
	commands := event_reaction.CreateCommands(conf)
	for event := range eventChannel {
		reaction, found := commands[event.Event]
		if !found {
			fmt.Printf("reaction not found for event: %v", event.Event)
		}
		reaction.Do(event)
	}
}
