package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"io"
	"log"
)

// logging will be implemented later
type Publisher struct {
	client  *redis.Client
	logger  io.Writer
	channel string
}

func CreatePublisher(url string, l io.Writer, channel string) *Publisher {
	c, err := connection.SetUpRedis(url)
	if err != nil {
		log.Fatal(fmt.Sprintf("can not create publisher: %v", err))
	}

	return &Publisher{
		client: c,
		logger: l,
	}
}

func (p *Publisher) NotifyUserSuccess(correlationId int, event string) {
	notifyEvent := domain.NotifyEvent{
		CorrelationId: correlationId,
		Event:         event,
	}

	eventBytes, err := json.Marshal(notifyEvent)
	if err != nil {
		_, _ = p.logger.Write([]byte(err.Error()))
	}

	p.client.Publish(context.Background(), p.channel, eventBytes)

}
