package event_reaction

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/query"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

func init() {
	Commands[domain.AuctionWinnerAnnounced] = WinnerAnnouncedCommand{}
}

type WinnerAnnouncedCommand struct{}

func (w WinnerAnnouncedCommand) Execute(application app.Application, event domain.Event) error {
	var winnerMessage domain.WinnerAnnounced
	if err := json.Unmarshal(event.Payload, &winnerMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling winner message: %v", err))
	}

	auctionState, readError := application.Queries.GetAuctionState.Handle(query.GetAuctionStateCommand{
		Ctx:       context.Background(),
		AuctionId: winnerMessage.GetAuctionId(),
	})
	if readError != nil {
		return errors.New(fmt.Sprintf("Error happened while reading auction state: %v", readError))
	}

	if err := application.Commands.UpdateState.Handle(
		command.UpdateStateCommand{
			CurrentState: auctionState,
			Event:        winnerMessage,
		}); err != nil {
		return errors.New(fmt.Sprintf("Error happened with saving winner: %v", err))
	}

	if err := application.Commands.SaveAuctionEvent.Handle(event.Event, winnerMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened during saving the auction event: %v", err))
	}
	return nil
}
