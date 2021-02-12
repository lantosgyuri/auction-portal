package domain

const AuctionRequested = "User_created_auction"
const AuctionWinnerAnnounced = "Auction_winner_announced"

type Event struct {
	Event   string
	Payload []byte
}
