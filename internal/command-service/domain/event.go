package domain

const (
	AuctionRequested       = "Create_auction"
	AuctionWinnerAnnounced = "Announce_winner"
	BidPlaceRequested      = "Place_bid"
	BidDeleteRequested     = "Delete_bid"
	UserCreateRequested    = "Create_user"
	UserDeleteRequested    = "Delete_user"
)

type Event struct {
	Event         string
	CorrelationId int
	Payload       []byte
}

type NotifyEvent struct {
	CorrelationId int
	Event         string
	Success       bool
	Error         string
}
