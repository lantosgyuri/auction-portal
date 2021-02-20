package domain

const AuctionRequested = "Create_auction"
const AuctionWinnerAnnounced = "Announce_winner"
const BidPlaceRequested = "Place_bid"
const BidDeleteRequested = "Delete_bid"
const UserCreateRequested = "Create_user"
const UserDeleteRequested = "Delete_user"

type Event struct {
	Event   string
	Payload []byte
}
