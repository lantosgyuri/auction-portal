package query

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type BidsForAuctionHandler struct {
	reader BidsForAuctionReader
}

type BidsForAuctionReader interface {
	FindBidsForAuction(ctx context.Context, auctionId string) ([]domain.AuctionEvent, error)
}

func (b BidsForAuctionHandler) Handle(ctx context.Context, auctionId string) ([]domain.AuctionEvent, error) {
	return b.reader.FindBidsForAuction(ctx, auctionId)
}
