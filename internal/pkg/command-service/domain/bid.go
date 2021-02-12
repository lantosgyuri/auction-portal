package domain

type BidEvent interface {
	GetUserId() int
}

type BidEventRaw struct {
	AuctionId string
	UserId    int
	EventType string
	BidId     int
	Amount    int
	TimeStamp int
}

type BidPlaced struct {
	AuctionId string
	Promoted  bool
	UserId    int
	Amount    int
	TimeStamp int
}

type BidDeleted struct {
	AuctionId  string
	ShouldSwap bool
	BidId      int
	UserId     int
	Amount     int
	TimeStamp  int
}

func (b BidPlaced) GetAuctionId() string {
	return b.AuctionId
}
func (b BidPlaced) GetUserId() int {
	return b.UserId
}

func (b BidDeleted) GetAuctionId() string {
	return b.AuctionId
}
func (b BidDeleted) GetUserId() int {
	return b.UserId
}
