package domain

const AuctionRequested = "User_created_auction"
const AuctionWinnerAnnounced = "Auction_winner_announced"
const BidPlaceRequested = "Bid_placed"
const BidDeleteRequested = "Bid_deleted"

type Event struct {
	Event   string
	Payload []byte
}
