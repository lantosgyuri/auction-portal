package domain

type BidPlaced struct {
	AuctionId int
	UserId    int
	Amount    int
	TimeStamp int
}

type BidDeleted struct {
	AuctionId int
	UserId    int
	TimeStamp int
}

type BidDoubled struct {
	AuctionId int
	UserId    int
	Amount    int
	TimeStamp int
}

func (b BidPlaced) isAuctionEvent()  {}
func (b BidDeleted) isAuctionEvent() {}
func (b BidDoubled) isAuctionEvent() {}
