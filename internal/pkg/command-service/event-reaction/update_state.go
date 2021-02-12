package event_reaction

import (
	"context"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

func UpdateState(ctx context.Context, message domain.AuctionEvent, application app.Application) error {
	auctionState, readError := application.Queries.GetAuctionState.Handle(ctx, message.GetAuctionId())
	if readError != nil {
		return errors.New(fmt.Sprintf("Error happened while reading auction state: %v", readError))
	}

	if err := application.Commands.UpdateState.Handle(
		command.UpdateStateCommand{
			CurrentState: auctionState,
			Event:        message,
		}); err != nil {
		return errors.New(fmt.Sprintf("Error happened with saving winner: %v", err))
	}

	return nil
}
