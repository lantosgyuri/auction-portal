package data_transformer

type Event struct {
	EventName string
	Payload   []byte
}

const (
	AuctionCreated  = "AuctionCreated"
	BidPlaced       = "BidPlaced"
	BidDeleted      = "BidDeleted"
	WinnerAnnounced = "WinnerAnnounced"
	UserCreated     = "UserCreated"
	UserDeleted     = "UserDeleted"
)

type AuctionCreatedEvent struct {
	ID        string
	Name      string
	StartDate int
	DueDate   int
}

type BidPlacedEvent struct {
	AuctionID string
	BidID     int
	UserID    int
	Amount    int
	Promote   bool
	Rollback  bool
}

type BidDeletedEvent struct {
	AuctionID string
	UserID    int
	BidID     int
	Amount    int
}

type WinnerAnnouncedEvent struct {
	UserID      int
	AuctionID   string
	AuctionName string
	UserName    string
}

type UserCreatedEvent struct {
	UserID   int
	Name     string
	Password string
}

type UserDeletedEvent struct {
	UserID int
}
