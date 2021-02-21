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

type WinnerAnnouncedEventHandler interface {
	Handle(context context.Context, event domain.WinnerAnnounced) error
}

type WinnerAnnouncedCommand struct {
	handler   WinnerAnnouncedEventHandler
	preserver AuctionEventPreserver
}

func CreateWinnerAnnouncedCommand() WinnerAnnouncedCommand {
	handler := command.AnnounceWinnerHandler{Repo: adapter.MariaDbStateRepository{Db: connection.SotDb}}
	preserver := command.SaveAuctionEventHandler{Repo: adapter.MariaDbAuctionRepository{Db: connection.SotDb}}
	return CreateWinnerCommandWithInterfaces(handler, preserver)
}

func CreateWinnerCommandWithInterfaces(handler WinnerAnnouncedEventHandler, preserver AuctionEventPreserver) WinnerAnnouncedCommand {
	return WinnerAnnouncedCommand{
		handler:   handler,
		preserver: preserver,
	}
}

func (w WinnerAnnouncedCommand) Execute(event domain.Event) error {
	var winnerMessage domain.WinnerAnnounced
	if err := json.Unmarshal(event.Payload, &winnerMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling winner message: %v", err))
	}

	if err := w.handler.Handle(context.Background(), winnerMessage); err != nil {
		return errors.New(fmt.Sprintf("Can not update state: %v", err))
	}

	if err := w.preserver.Handle(event.Event, winnerMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened during saving the auction event: %v", err))
	}
	return nil
}
