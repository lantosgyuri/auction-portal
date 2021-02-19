package adapter

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"gorm.io/gorm"
)

type MariaDbStateRepository struct {
	Db *gorm.DB
}

func (m MariaDbStateRepository) UpdateState(
	ctx context.Context,
	event domain.AuctionEvent,
	update func(auction domain.Auction) (domain.Auction, error)) error {

	var currentState domain.Auction

	m.Db.First(&currentState, "Id = ?", event.GetAuctionId())

	newState, err := update(currentState)

	if err != nil {
		return err
	}

	return m.Db.Model(&newState).Updates(newState).Error
}
