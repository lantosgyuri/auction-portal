package command

import (
	"github.com/lantosgyuri/auction-portal/internal/data-transformer/event-reaction"
)

type AuctionRepository interface {
	SaveAuction(auction event_reaction.AuctionCreatedEvent) error
	UpdateAuctionBid(bid event_reaction.BidPlacedEvent) error
	UpdateAuctionWinner(winner event_reaction.WinnerAnnouncedEvent) error
}

type BidRepository interface {
	CreateUser(user event_reaction.UserCreatedEvent) error
	SaveUserBid(bid event_reaction.BidPlacedEvent) error
	DeleteUserBid(bid event_reaction.BidDeletedEvent) error
	DeleteUserEntries(user event_reaction.UserDeletedEvent) error
}

type UserRepository interface {
	SaveUser(user event_reaction.UserCreatedEvent) error
	DeleteUser(user event_reaction.UserDeletedEvent) error
}

type WinnerRepository interface {
	SaveWinner(winner event_reaction.WinnerAnnouncedEvent) error
}

/*
Domain:
List auctions. (Auction table)
List winners. (Winner table)
Get all bid from one user for an auction. (User table)
If a bid is deleted and have to roll back, than there should be also a place bid message with the rollback.
Delete user bid from an auction. (User table)

Tables:
Auction: Id, Name, StartDate, EndDate, CurrentPrice(maybe null), CurrentUser(maybe null), Winner(maybe null), Promoted
User: Id, Name
Bid: UserId, auctionId, placedBids[], deletedBids[]
		bid{ id, amount, type(place, delete, promote) }
Winner: userId, auctionId, userName, auctionName
*/
