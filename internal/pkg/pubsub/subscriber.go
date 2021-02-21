package pubsub

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
)

type Subscriber struct {
	client         *redis.Client
	channels       []string
	subscriberType interface{}
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

func (s *Subscriber) AddChannel(ch string) {
	s.channels = append(s.channels, ch)
}

func (s *Subscriber) Get(eventChan chan []byte) {
	for _, v := range s.channels {
		pubs := s.client.Subscribe(context.Background(), v)
		ch := pubs.Channel()
		go consumeEvents(ch, eventChan)
	}
}

func consumeEvents(ch <-chan *redis.Message, eventChan chan []byte) {
	for msg := range ch {
		eventChan <- []byte(msg.Payload)
	}
}
