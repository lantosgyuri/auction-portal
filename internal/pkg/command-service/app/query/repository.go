package query

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type AuctionStateReader interface {
	FindAuction(ctx context.Context, auctionId string) (domain.Auction, error)
}
