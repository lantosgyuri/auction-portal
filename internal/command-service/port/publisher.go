package port

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"log"
)

const (
	BidChannel     = "bid"
	AuctionChannel = "auction"
	UserChannel    = "user"
)

// logging will be implemented later
type Logger interface {
	Log(s string)
}

type FakeLogger struct{}

func (l FakeLogger) Log(s string) {
	fmt.Println(s)
}

type publisher struct {
	client  *redis.Client
	logger  Logger
	channel string
}

func CreatePublisher(url string, l Logger, ch string) *publisher {
	c, err := connection.SetUpRedis(url)
	if err != nil {
		log.Fatal(fmt.Sprintf("can not create publisher: %v", err))
	}

	return &publisher{
		client:  c,
		logger:  l,
		channel: ch,
	}
}

// Currently it adds only a prefix, but in a real prod env, this should be a separate struct wit ha separate connection
func (p *publisher) NotifyUserSuccess(event domain.NotifyEvent) {
	event.Success = true
	eventBytes, err := json.Marshal(event)
	if err != nil {
		p.logger.Log(err.Error())
	}

	p.client.Publish(context.Background(), addPrefix(p.channel), eventBytes)

}

// Currently it adds only a prefix, but in a real prod env, this should be a separate struct wit ha separate connection
func (p *publisher) NotifyUserFail(event domain.NotifyEvent) {
	event.Success = false
	eventBytes, err := json.Marshal(event)
	if err != nil {
		p.logger.Log(err.Error())
	}
	p.client.Publish(context.Background(), addPrefix(p.channel), eventBytes)
}

func (p *publisher) PublishData(event domain.Event) {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		p.logger.Log(err.Error())
	}

	p.client.Publish(context.Background(), p.channel, eventBytes)

}

func addPrefix(s string) string {
	return "Notify" + s
}
