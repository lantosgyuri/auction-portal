package command

import "github.com/lantosgyuri/auction-portal/internal/command-service/domain"

type AuctionRepository interface {
	SaveAuction(auction domain.CreateAuctionRequested) error
	SaveWinner(winner domain.WinnerAnnounced) error
}

type BidRepository interface {
	SaveBid(bid domain.BidPlaced) error
	DeleteBid(bid domain.BidDeleted) error
}

type UserRepository interface {
	SaveUser(user domain.CreateUserRequested) error
	DeleteUser(user domain.DeleteUserRequest) error
}
