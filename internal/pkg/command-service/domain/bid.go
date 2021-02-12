package domain

type BidEvent interface {
	GetUserId() int
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
