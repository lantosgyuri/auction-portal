package command

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type AuctionRepository interface {
	SaveAuctionEvent(event domain.AuctionEventRaw) error
	CreateNewAuction(auction domain.Auction) error
}

type StateRepository interface {
	UpdateState(
		context context.Context,
		event domain.AuctionEvent,
		update func(auction domain.Auction) domain.Auction,
	) error
}

type BidRepository interface {
	SaveBidEvent(event domain.BidEventRaw) error
	CreateBid(bid domain.BidPlaced) error
	DeleteBid(bid domain.BidDeleted) error
}

type UserRepository interface {
	SaveUserEvent() error
}
