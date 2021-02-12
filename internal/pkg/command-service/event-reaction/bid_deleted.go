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
	Commands[domain.BidDeleteRequested] = BidDeleteRequestedCommand{}
}

type BidDeleteRequestedCommand struct{}

func (b BidDeleteRequestedCommand) Execute(application app.Application, event domain.Event) error {
	var bidDeleteMessage domain.BidDeleted

	if err := json.Unmarshal(event.Payload, &bidDeleteMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling winner message: %v", err))
	}

	auctionState, readError := application.Queries.GetAuctionState.Handle(query.GetAuctionStateCommand{
		Ctx:       context.Background(),
		AuctionId: bidDeleteMessage.GetAuctionId(),
	})
	if readError != nil {
		return errors.New(fmt.Sprintf("Error happened while reading auction state: %v", readError))
	}

	if auctionState.CurrentUser == bidDeleteMessage.UserId && auctionState.CurrentBid == bidDeleteMessage.Amount {
		bids, err := application.Queries.GetBidsForAuction.Handle(query.BidsForAuctionCommand{
			Ctx:       context.Background(),
			AuctionId: bidDeleteMessage.GetAuctionId(),
		})
		if err != nil {
			return errors.New(fmt.Sprintf("Error happened while reading bids: %v", err))
		}

		fallbackToSecondHighest(bids, &bidDeleteMessage)
	}

	if err := application.Commands.UpdateState.Handle(
		command.UpdateStateCommand{
			CurrentState: auctionState,
			Event:        bidDeleteMessage,
		}); err != nil {
		return errors.New(fmt.Sprintf("Error happened with saving winner: %v", err))
	}

	if err := application.Commands.DeleteBid.Handle(bidDeleteMessage); err != nil {
		return err
	}
	if err := application.Commands.SaveBidEvent.Handle(event.Event, bidDeleteMessage); err != nil {
		return err
	}

	return nil
}

func fallbackToSecondHighest(bids []domain.Bid, bidDeletedMessage *domain.BidDeleted) {
	bidDeletedMessage.ShouldSwap = true
	if len(bids) > 1 {
		bidDeletedMessage.UserId = bids[1].UserId
		bidDeletedMessage.Amount = bids[1].Amount
		return
	}
	bidDeletedMessage.Amount = 0
	bidDeletedMessage.UserId = 0
	return
}
