package event_reaction

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/adapter"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
)

type BidPLacedEventHandler interface {
	Handle(ctx context.Context, placed domain.BidPlaced) error
}

type PreserveBidEvent interface {
	Handle(eventName string, event domain.BidEvent) error
}

type BidPlaceRequestedCommand struct {
	handler   BidPLacedEventHandler
	preserver PreserveBidEvent
	publisher EventPublisher
}

func CreateBidPlacedReqCommand() BidPlaceRequestedCommand {
	handler := command.PlaceBidHandler{
		StateRepo: adapter.MariaDbStateRepository{
			Db: connection.SotDb,
		},
		BidRepo: adapter.MariaDbBidRepository{
			Db: connection.SotDb},
	}
	preserver := command.SaveBidEventHandler{
		Repo: adapter.MariaDbBidRepository{
			Db: connection.SotDb},
	}

	return CreateBidPlacedReqWithInterfaces(handler, preserver)
}

func CreateBidPlacedReqWithInterfaces(handler BidPLacedEventHandler, preserver PreserveBidEvent) BidPlaceRequestedCommand {
	return BidPlaceRequestedCommand{
		handler:   handler,
		preserver: preserver,
	}
}

func (b BidPlaceRequestedCommand) Execute(event domain.Event) error {
	var bidPlacedMessage domain.BidPlaced
	if err := json.Unmarshal(event.Payload, &bidPlacedMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling winner message: %v", err))
	}
	if err := b.handler.Handle(context.Background(), bidPlacedMessage); err != nil {
		return err
	}
	if err := b.preserver.Handle(event.Event, bidPlacedMessage); err != nil {
		return err
	}
	if err := b.publisher.Publish(event); err != nil {
		return errors.New(fmt.Sprintf("Can not publish event: %v", err))
	}

	return nil
}
