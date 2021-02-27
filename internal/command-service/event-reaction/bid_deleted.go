package event_reaction

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/adapter"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/config"
)

type BidDeletedEventHandler interface {
	Handle(ctx context.Context, event domain.BidDeleted) error
}

type BidDeleteRequestedCommand struct {
	handler   BidDeletedEventHandler
	preserver PreserveBidEvent
	publisher EventPublisher
}

func CreateBidDeletedCommand(conf config.CommandService) BidDeleteRequestedCommand {
	handler := command.DeleteBidHandler{
		BidRepo:   adapter.CreateMariaDbBidRepository(),
		StateRepo: adapter.CreateMariaDbStateRepository(),
	}
	preserver := command.SaveBidEventHandler{
		Repo: adapter.CreateMariaDbBidRepository(),
	}

	return CreateBidDeletedWithInterfaces(handler, preserver)
}

func CreateBidDeletedWithInterfaces(handler BidDeletedEventHandler, preserver PreserveBidEvent) BidDeleteRequestedCommand {
	return BidDeleteRequestedCommand{
		handler:   handler,
		preserver: preserver,
	}
}

func (b BidDeleteRequestedCommand) Execute(event domain.Event) {
	var bidDeleteMessage domain.BidDeleted

	if err := json.Unmarshal(event.Payload, &bidDeleteMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling winner message: %v", err))
	}

	if err := b.handler.Handle(context.Background(), bidDeleteMessage); err != nil {
		return err
	}
	if err := b.preserver.Handle(event.Event, bidDeleteMessage); err != nil {
		return err
	}
	if err := b.publisher.Publish(event); err != nil {
		return errors.New(fmt.Sprintf("Can not publish event: %v", err))
	}
	return nil
}
