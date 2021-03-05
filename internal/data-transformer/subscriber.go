package data_transformer

import (
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
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
	for event := range eventChannel {
		fmt.Printf("I got events %v", event)
	}
}
