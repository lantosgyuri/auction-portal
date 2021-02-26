package adapter

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"gorm.io/gorm"
)

type MariaDbStateRepository struct {
	db *gorm.DB
}

func CreateMariaDbStateRepository() MariaDbStateRepository {
	return MariaDbStateRepository{
		db: connection.GetMariDbConnection(),
	}
}

func (m MariaDbStateRepository) UpdateState(
	ctx context.Context,
	event domain.AuctionEvent,
	update func(auction domain.Auction) (domain.Auction, error)) error {

	var currentState domain.Auction

	m.db.First(&currentState, "Id = ?", event.GetAuctionId())

	newState, err := update(currentState)

	if err != nil {
		return err
	}

	state := make(map[string]interface{})

	state["winner"] = newState.Winner
	state["current_bid"] = newState.CurrentBid
	state["current_user"] = newState.CurrentUser
	state["promoted"] = newState.Promoted
	state["place_event_count"] = newState.PlaceEventCount
	state["delete_event_count"] = newState.DeleteEventCount

	return m.db.Model(&newState).Updates(state).Error
}
