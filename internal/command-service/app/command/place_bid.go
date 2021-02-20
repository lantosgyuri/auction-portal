package command

import (
	"context"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"time"
)

type PlaceBidHandler struct {
	BidRepo   BidRepository
	StateRepo StateRepository
}

func (p PlaceBidHandler) Handle(ctx context.Context, bid domain.BidPlaced) error {

	highestBid := p.BidRepo.IsHighestUserBid(ctx, bid, func(highestBid domain.Bid) bool {
		return bid.Amount > highestBid.Amount
	})

	if !highestBid {
		return errors.New("bid is smaller than latest bid")
	}

	b := domain.Bid{
		UserId:    bid.UserId,
		Amount:    bid.Amount,
		Promoted:  bid.Promoted,
		AuctionId: bid.AuctionId,
	}

	if err := p.BidRepo.SaveBid(b); err != nil {
		return errors.New(fmt.Sprintf("can not save bid: %v", err))
	}

	return p.StateRepo.UpdateState(ctx, bid, func(auction domain.Auction) (domain.Auction, error) {

		if auction.CurrentBid >= bid.Amount {
			return domain.Auction{}, errors.New("ignore bid. Bid is smaller or equal with current bid")
		}

		if auction.DueDate <= int(time.Now().Unix()) {
			return domain.Auction{}, errors.New("auction due date is in the past")
		}

		return domain.ApplyOnSnapshot(auction, bid), nil
	})
}
