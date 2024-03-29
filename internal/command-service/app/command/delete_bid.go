package command

import (
	"context"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

type DeleteBidHandler struct {
	BidRepo   BidRepository
	StateRepo StateRepository
}

func (d DeleteBidHandler) Handle(ctx context.Context, deletedBid domain.BidDeleted) error {
	if err := d.BidRepo.IsHighestAuctionBid(
		ctx,
		deletedBid.GetAuctionId(),
		func(topBids []domain.Bid) error {
			if len(topBids) == 0 {
				return errors.New(fmt.Sprintf("no bids for auction %v", deletedBid.AuctionId))
			}

			if isHighestBid(topBids[0], deletedBid) {
				deletedBid.ShouldSwap = true
				deletedBid.UserId = 0
				deletedBid.Amount = 0
				if canFallbackToBid(topBids) {
					deletedBid.UserId = topBids[1].UserId
					deletedBid.Amount = topBids[1].Amount
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

	bid := domain.Bid{
		AuctionId: deletedBid.AuctionId,
		UserId:    deletedBid.UserId,
		Amount:    deletedBid.Amount,
		Id:        deletedBid.BidId,
	}

	return d.BidRepo.DeleteBid(bid)
}

func isHighestBid(topBid domain.Bid, deletedBid domain.BidDeleted) bool {
	return topBid.Id == deletedBid.BidId
}

func canFallbackToBid(bids []domain.Bid) bool {
	return len(bids) > 1
}
