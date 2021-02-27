package port

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/pubsub"
)

type Channel int

const (
	Auction Channel = 1 << iota
	Bid
	User
	NotifyAuction
	NotifyBid
	NotifyUser
)

var channelMapping = map[Channel]string{
	Auction: "Auction",
	Bid:     "Bid",
	User:    "User",
}

type EventSender struct {
	publisher pubsub.Publisher
	channel   Channel
}

func CreateEventSender(publisher pubsub.Publisher, channel Channel) EventSender {
	return EventSender{
		publisher: publisher,
		channel:   channel,
	}
}

func (e EventSender) Publish(event domain.Event) error {
	return e.publisher.SendEvent(event, channelMapping[e.channel], event.Event)
}
