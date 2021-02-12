package query

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type GetAuctionStateCommand struct {
	Ctx       context.Context
	AuctionId string
}

type AuctionStateHandler struct {
	Reader AuctionStateReader
}

func (a AuctionStateHandler) Handle(cmd GetAuctionStateCommand) (domain.Auction, error) {
	return a.Reader.FindAuction(cmd.Ctx, cmd.AuctionId)
}
