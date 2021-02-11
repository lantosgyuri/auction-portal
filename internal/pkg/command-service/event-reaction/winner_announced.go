package event_reaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

func init() {
	Commands[domain.AuctionWinnerAnnounced] = WinnerAnnouncedCommand{}
}

type WinnerAnnouncedCommand struct{}

func (w WinnerAnnouncedCommand) Execute(application app.Application, event domain.Event) error {
	var winnerMessage domain.AuctionWinnerMessage
	if err := json.Unmarshal(event.Payload, &winnerMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling winner message: %v", err))
	}
	if err := application.Commands.SaveWinner.Handle(winnerMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened with saving winner: %v", err))
	}
	if err := application.Commands.SaveAuctionEvent.Handle(event); err != nil {
		return errors.New(fmt.Sprintf("Error happened during saving the auction event: %v", err))
	}
	return nil
}
