package adapter

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"gorm.io/gorm"
)

type MariaDbAuctionRepository struct {
	db *gorm.DB
}

func CreateMariaDbAuctionRepository() MariaDbAuctionRepository {
	return MariaDbAuctionRepository{
		db: connection.GetMariDbConnection(),
	}
}

func (m MariaDbAuctionRepository) SaveAuctionEvent(event domain.AuctionEventRaw) error {
	return m.db.Create(&event).Error
}

func (m MariaDbAuctionRepository) CreateNewAuction(auction domain.Auction) error {
	return m.db.Create(&auction).Error
}
