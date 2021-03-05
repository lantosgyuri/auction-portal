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

/*
Domain:
List auctions. (Auction table)
List winners. (Winner table)
Get all bid from one user for an auction. (User table)
If a bid is deleted and have to roll back, than there should be also a place bid message with the rollback.
Delete user bid from an auction. (User table)

Tables:
Auction: Id, Name, StartDate, EndDate, CurrentPrice(maybe null), CurrentUser(maybe null), Winner(maybe null)
User: Id, Name
Bid: UserId, auctionId, placedBids[], deletedBids[]
		bid{ id, amount, type }
Winner: userId, auctionId, userName, auctionName
*/
