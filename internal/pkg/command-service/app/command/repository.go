package command

import "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"

type AuctionRepository interface {
	SaveAuctionEvent(event domain.AuctionEventRaw) error
	UpdateAuctionState(event domain.Auction) error
	CreateNewAuction(auction domain.Auction) error
}

type BidRepository interface {
	SaveBidEvent(event domain.BidEventRaw) error
	CreateBid(bid domain.BidPlaced) error
	DeleteBid(bid domain.BidDeleted) error
}

type UserRepository interface {
	SaveUserEvent() error
}
