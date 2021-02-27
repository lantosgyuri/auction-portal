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

type Publisher struct {
	client *redis.Client
}

func CreatePublisher(url string) (*Publisher, error) {
	c, err := connection.SetUpRedis(url)
	if err != nil {
		return nil, errors.New("can not create redis connection")
	}

	return &Publisher{
		client: c,
	}, nil
}

func (p *Publisher) SendEvent(message interface{}, channel string, eventName string) error {
	messageBytes, err := json.Marshal(message)

	if err != nil {
		return errors.New(fmt.Sprintf("can not marshal message: %v", err))
	}

	event := domain.Event{
		Event:   eventName,
		Payload: messageBytes,
	}

	eventBytes, err := json.Marshal(event)

	if err != nil {
		return errors.New(fmt.Sprintf("can not marshal event: %v", err))
	}

	p.client.Publish(context.Background(), channel, eventBytes)
	return nil
}
