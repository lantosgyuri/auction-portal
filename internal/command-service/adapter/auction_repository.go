package adapter

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"gorm.io/gorm"
)

type MariaDbAuctionRepository struct {
	Db *gorm.DB
}

func (m MariaDbAuctionRepository) SaveAuctionEvent(event domain.AuctionEventRaw) error {
	return m.Db.Create(&event).Error
}

func (m MariaDbAuctionRepository) CreateNewAuction(auction domain.Auction) error {
	return m.Db.Create(&auction).Error
}
