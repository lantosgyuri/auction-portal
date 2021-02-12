package query

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type BidsForAuctionCommand struct {
	Ctx       context.Context
	AuctionId string
}

type BidForAuctionHandler struct {
	Reader BidReader
}

func (b BidForAuctionHandler) Handle(cmd BidsForAuctionCommand) ([]domain.Bid, error) {
	return b.Reader.FindBidsForAuctionDescending(cmd.Ctx, cmd.AuctionId)
}
