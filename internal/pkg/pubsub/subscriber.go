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

type Subscriber struct {
	client *redis.Client
}

func CreateSubscriber(url string) (*Subscriber, error) {
	c, err := connection.SetUpRedis(url)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("can not create redis connection: %v", err))
	}

	return &Subscriber{
		client: c,
	}, nil
}

func (s *Subscriber) GetEvents(channel string, eventChan chan domain.Event) {
	pubs := s.client.Subscribe(context.Background(), channel)

	ch := pubs.Channel()

	for msg := range ch {
		var event domain.Event
		if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
			fmt.Printf("error happened during unmarshal event: %v", err)
		}
		eventChan <- event
	}

}
