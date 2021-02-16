package adapter

import (
	"context"
	"errors"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"gorm.io/gorm"
)

type MariaDbBidRepository struct {
	Db *gorm.DB
}

func (m MariaDbBidRepository) SaveBidEvent(event domain.BidEventRaw) error {
	return m.Db.Create(&event).Error
}

func (m MariaDbBidRepository) SaveBid(bid domain.Bid) error {
	return m.Db.Create(&bid).Error
}

func (m MariaDbBidRepository) DeleteBid(bid domain.Bid) error {
	return m.Db.Delete(&bid).Error
}

func (m MariaDbBidRepository) IsHighestUserBid(ctx context.Context, placed domain.BidPlaced,
	validate func(userHighestBid domain.Bid) bool) bool {
	var bid domain.Bid
	result := m.Db.Where(&domain.Bid{UserId: placed.UserId, AuctionId: placed.AuctionId}).First(&bid)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}

	return validate(bid)
}

func (m MariaDbBidRepository) IsHighestAuctionBid(ctx context.Context, auctionId string,
	onHighestBid func(topBid domain.Bid, secondBid domain.Bid) error) error {
	bids := make([]domain.Bid, 2)

	m.Db.Where("AuctionId = ?", auctionId).Limit(2).First(&bids)

	return onHighestBid(bids[0], bids[1])
}
