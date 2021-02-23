package adapter

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
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

	state := make(map[string]interface{})

	state["winner"] = newState.Winner
	state["current_bid"] = newState.CurrentBid
	state["current_user"] = newState.CurrentUser
	state["promoted"] = newState.Promoted
	state["place_event_count"] = newState.PlaceEventCount
	state["delete_event_count"] = newState.DeleteEventCount

	return m.Db.Model(&newState).Updates(state).Error
}
