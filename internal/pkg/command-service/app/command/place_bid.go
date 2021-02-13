package command

import (
	"context"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
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

	if err := p.BidRepo.SaveBid(bid); err != nil {
		return errors.New(fmt.Sprintf("can not save bid: %v", err))
	}

	return p.StateRepo.UpdateState(ctx, bid, func(auction domain.Auction) (domain.Auction, error) {

		if auction.CurrentBid >= bid.Amount {
			return domain.Auction{}, errors.New("ignore bid. Bid is smaller or equal with current bid")
		}

		return domain.ApplyOnSnapshot(auction, bid), nil
	})
}
