package command

import "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"

type AuctionRepository interface {
	SaveAuctionEvent(event domain.AuctionEventRaw) error
	UpdateAuctionState(event domain.Auction) error
	CreateNewAuction(auction domain.Auction) error
}

type BidRepository interface {
	SaveBidEvent(event domain.BidEventRaw) error
}

type UserRepository interface {
	SaveUSerEvent() error
}
