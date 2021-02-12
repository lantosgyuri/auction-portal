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
	Commands[domain.BidPlaceRequested] = BidPlaceRequestedCommand{}
}

type BidPlaceRequestedCommand struct{}

func (b BidPlaceRequestedCommand) Execute(application app.Application, event domain.Event) error {
	var bidPlacedMessage domain.BidPlaced
	if err := json.Unmarshal(event.Payload, &bidPlacedMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling winner message: %v", err))
	}

	highestBid, readErrorBids := application.Queries.GetHighestBidForAuction.Handle(query.UserHighestBidCommand{
		UserId:    bidPlacedMessage.UserId,
		AuctionId: bidPlacedMessage.GetAuctionId(),
		Ctx:       context.Background(),
	})

	if readErrorBids != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling winner message: %v", readErrorBids))
	}

	if highestBid.Amount > bidPlacedMessage.Amount {
		return errors.New("Bid is smaller than bid before")
	}

	auctionState, readErrorAuctionState := application.Queries.GetAuctionState.Handle(query.GetAuctionStateCommand{
		Ctx:       context.Background(),
		AuctionId: bidPlacedMessage.GetAuctionId(),
	})

	if readErrorAuctionState != nil {
		return errors.New(fmt.Sprintf("Error happened while reading auction state: %v", readErrorAuctionState))
	}

	if auctionState.CurrentBid < bidPlacedMessage.Amount {
		if err := application.Commands.UpdateState.Handle(
			command.UpdateStateCommand{
				CurrentState: auctionState,
				Event:        bidPlacedMessage,
			}); err != nil {
			return errors.New(fmt.Sprintf("Error happened with saving winner: %v", err))
		}
	}

	if err := application.Commands.CreateBid.Handle(bidPlacedMessage); err != nil {
		return err
	}
	if err := application.Commands.SaveBidEvent.Handle(event.Event, bidPlacedMessage); err != nil {
		return err
	}

	return nil
}
