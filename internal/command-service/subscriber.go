package command_service

import (
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/command-service/event-reaction"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
	"github.com/lantosgyuri/auction-portal/internal/pkg/pubsub"
	"sync"
)

func StartSubscriber(conf config.CommandService, parentWg *sync.WaitGroup) {
	eventChannel := make(chan domain.Event)
	eventSubscriber, err := pubsub.CreateSubscriber(conf.RedisConf.WriteUrl)

	if err != nil {
		fmt.Printf("error happened during create subscriber: %v", err)
		parentWg.Done()
	}

	go pubsub.SubscribeToMainEvents(eventSubscriber, eventChannel)
	go consumeMessages(conf, eventChannel)
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
