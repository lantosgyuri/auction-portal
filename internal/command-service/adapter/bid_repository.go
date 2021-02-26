package adapter

import (
	"context"
	"errors"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"gorm.io/gorm"
)

type MariaDbBidRepository struct {
	db *gorm.DB
}

func CreateMariaDbBidRepository() MariaDbBidRepository {
	return MariaDbBidRepository{
		db: connection.GetMariDbConnection(),
	}
}

func (m MariaDbBidRepository) SaveBidEvent(event domain.BidEventRaw) error {
	return m.db.Create(&event).Error
}

func (m MariaDbBidRepository) SaveBid(bid domain.Bid) error {
	return m.db.Create(&bid).Error
}

func (m MariaDbBidRepository) DeleteBid(bid domain.Bid) error {
	return m.db.Delete(&bid).Error
}

func (m MariaDbBidRepository) IsHighestUserBid(ctx context.Context, placed domain.BidPlaced,
	validate func(userHighestBid domain.Bid) bool) bool {
	var bid domain.Bid
	result := m.db.Where(&domain.Bid{UserId: placed.UserId, AuctionId: placed.AuctionId}).Last(&bid)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return true
	}

	return validate(bid)
}

func (m MariaDbBidRepository) IsHighestAuctionBid(ctx context.Context, auctionId string,
	onHighestBid func(topBids []domain.Bid) error) error {
	bids := make([]domain.Bid, 0)

	m.db.Where("auction_id = ?", auctionId).Order("amount desc").Limit(2).Find(&bids)

	return onHighestBid(bids)
}
