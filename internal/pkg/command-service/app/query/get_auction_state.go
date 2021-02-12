package query

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type AuctionStateHandler struct {
	Reader AuctionStateReader
}

func (a AuctionStateHandler) Handle(ctx context.Context, auctionId string) (domain.Auction, error) {
	return a.Reader.FindAuction(ctx, auctionId)
}
