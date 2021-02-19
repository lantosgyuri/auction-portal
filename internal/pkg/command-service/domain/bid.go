package domain

import "gorm.io/gorm"

type BidEvent interface {
	GetUserId() int
}

type Bid struct {
	gorm.Model
	Id        int `gorm:"primaryKey"`
	UserId    int
	AuctionId string `gorm:"size:191"`
	Promoted  bool
	Amount    int
	Auction   Auction
	User      User
}

type BidEventRaw struct {
	gorm.Model
	Id        int `gorm:"primaryKey"`
	AuctionId string
	UserId    int
	EventType string
	BidId     int
	Amount    int
}

type BidPlaced struct {
	AuctionId string
	Promoted  bool
	UserId    int
	Amount    int
}

type BidDeleted struct {
	AuctionId  string
	ShouldSwap bool
	BidId      int
	UserId     int
	Amount     int
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
