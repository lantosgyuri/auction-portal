package query

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type UserHighestBidCommand struct {
	Ctx       context.Context
	UserId    int
	AuctionId string
}

type UserHighestBidHandler struct {
	Reader BidReader
}

func (g UserHighestBidHandler) Handle(cmd UserHighestBidCommand) (domain.Bid, error) {
	return g.Reader.FindUserHighestBidForAuction(cmd.Ctx, cmd.AuctionId, cmd.UserId)
}
