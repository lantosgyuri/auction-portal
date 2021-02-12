package query

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type AuctionStateReader interface {
	FindAuction(ctx context.Context, auctionId string) (domain.Auction, error)
}

type BidReader interface {
	FindBidsForAuctionDescending(ctx context.Context, auctionId string) ([]domain.Bid, error)
	FindUserHighestBidForAuction(ctx context.Context, auctionID string, userId int) (domain.Bid, error)
}
