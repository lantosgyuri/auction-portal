package adapter

import (
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"gorm.io/gorm"
)

type MariaDbAuctionRepository struct {
	Db *gorm.DB
}

func (m MariaDbAuctionRepository) SaveAuctionEvent(event domain.AuctionEventRaw) error {
	m.Db.Create(&event)
	return nil
}

func (m MariaDbAuctionRepository) CreateNewAuction(auction domain.Auction) error {
	m.Db.Create(&auction)
	return nil
}
