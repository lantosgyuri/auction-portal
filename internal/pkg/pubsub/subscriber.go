package pubsub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
)

type EventSubscriber struct {
	client   *redis.Client
	channels []string
}

func CreateSubscriber(url string) (*EventSubscriber, error) {
	c, err := connection.SetUpRedis(url)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("can not create redis connection: %v", err))
	}

	return &EventSubscriber{
		client: c,
	}, nil
}

func (s *EventSubscriber) AddChannel(ch string) {
	s.channels = append(s.channels, ch)
}

func (s *EventSubscriber) Get(eventChan chan domain.Event) {
	for _, v := range s.channels {
		pubs := s.client.Subscribe(context.Background(), v)
		ch := pubs.Channel()
		go consumeEvents(ch, eventChan)
	}
}

func consumeEvents(ch <-chan *redis.Message, eventChan chan domain.Event) {
	for msg := range ch {
		var event domain.Event
		if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
			fmt.Printf("error happened during unmarshal event: %v", err)
		}
		eventChan <- event
	}
}
