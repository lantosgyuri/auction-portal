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

type BidDeletedEventHandler interface {
	Handle(ctx context.Context, event domain.BidDeleted) error
}

type BidDeleteRequestedCommand struct {
	handler   BidDeletedEventHandler
	preserver PreserveBidEvent
	publisher EventPublisher
}

func CreateBidDeletedCommand() BidDeleteRequestedCommand {
	handler := command.DeleteBidHandler{
		BidRepo:   adapter.MariaDbBidRepository{Db: connection.SotDb},
		StateRepo: adapter.MariaDbStateRepository{Db: connection.SotDb},
	}
	preserver := command.SaveBidEventHandler{
		Repo: adapter.MariaDbBidRepository{Db: connection.SotDb},
	}

	return CreateBidDeletedWithInterfaces(handler, preserver)
}

func CreateBidDeletedWithInterfaces(handler BidDeletedEventHandler, preserver PreserveBidEvent) BidDeleteRequestedCommand {
	return BidDeleteRequestedCommand{
		handler:   handler,
		preserver: preserver,
	}
}

func (b BidDeleteRequestedCommand) Execute(event domain.Event) error {
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
