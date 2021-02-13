package command

import (
	"context"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type DeleteBidHandler struct {
	BidRepo   BidRepository
	StateRepo StateRepository
}

func (d DeleteBidHandler) Handle(ctx context.Context, deletedBid domain.BidDeleted) error {
	if err := d.BidRepo.IsHighestAuctionBid(
		ctx,
		deletedBid.GetAuctionId(),
		func(topBid domain.Bid, secondBid domain.Bid) error {
			if isHighestBidFromUser(topBid, deletedBid) {
				deletedBid.ShouldSwap = true
				deletedBid.UserId = 0
				deletedBid.Amount = 0
				if canFallbackToBid(secondBid) {
					deletedBid.UserId = secondBid.UserId
					deletedBid.Amount = secondBid.Amount
				}
			}

			if updateErr := d.StateRepo.UpdateState(ctx, deletedBid, func(auction domain.Auction) (domain.Auction, error) {
				return domain.ApplyOnSnapshot(auction, deletedBid), nil
			}); updateErr != nil {
				return updateErr
			}
			return nil
		}); err != nil {
		return errors.New(fmt.Sprintf("can not validate delete request:  %v", err))
	}

	return d.BidRepo.DeleteBid(deletedBid)
}

func isHighestBidFromUser(topBid domain.Bid, deletedBid domain.BidDeleted) bool {
	return topBid.UserId == deletedBid.UserId && topBid.Amount == deletedBid.Amount
}

func canFallbackToBid(secondBid domain.Bid) bool {
	return secondBid.UserId != 0
}
