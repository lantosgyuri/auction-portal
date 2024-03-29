package command

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

type AuctionRepository interface {
	SaveAuctionEvent(event domain.AuctionEventRaw) error
	CreateNewAuction(auction domain.Auction) error
}

type StateRepository interface {
	UpdateState(
		context context.Context,
		event domain.AuctionEvent,
		update func(auction domain.Auction) (domain.Auction, error),
	) error
}

type BidRepository interface {
	SaveBidEvent(event domain.BidEventRaw) error
	IsHighestUserBid(
		context context.Context,
		bid domain.BidPlaced,
		validate func(userHighestBid domain.Bid) bool,
	) bool
	IsHighestAuctionBid(
		ctx context.Context,
		auctionId string,
		onHighestBid func(bids []domain.Bid) error,
	) error
	SaveBid(bid domain.Bid) error
	DeleteBid(bid domain.Bid) error
}

type UserRepository interface {
	SaveUserEvent(event domain.UserEventRaw) error
	CreateUser(user domain.User) error
	DeleteUser(user domain.User) error
}
