package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var messageCreate = CreateAuctionRequested{
	UUID:      "tttt-ttt",
	Name:      "Test",
	StartDate: 1,
	DueDate:   2,
}

var snapshot = Auction{
	UUID:        "ttt-ttt",
	Name:        "Test",
	StartDate:   1,
	DueDate:     2,
	CurrentBid:  0,
	CurrentUser: 0,
}

func TestNewAuction(t *testing.T) {
	auction := NewAuction(messageCreate)
	assert.EqualValues(t, "tttt-ttt", auction.UUID)
	assert.EqualValues(t, messageCreate.Name, auction.Name)
	assert.EqualValues(t, messageCreate.DueDate, auction.DueDate)
	assert.EqualValues(t, messageCreate.StartDate, auction.StartDate)
	assert.EqualValues(t, 0, auction.Winner)
}

func TestPlaceBid(t *testing.T) {
	placeBid := BidPlaced{
		Amount:   6,
		UserId:   1,
		Promoted: false,
	}

	newState := ApplyOnSnapshot(snapshot, placeBid)

	assert.EqualValues(t, 6, newState.CurrentBid)
	assert.EqualValues(t, 1, newState.CurrentUser)
	assert.EqualValues(t, false, newState.Promoted)
	assert.EqualValues(t, 1, newState.PlaceEventCount)
	assert.EqualValues(t, "Test", newState.Name)

	placeBidSecond := BidPlaced{
		Amount: 6,
		UserId: 1,
	}
	newStateSecond := ApplyOnSnapshot(newState, placeBidSecond)

	assert.EqualValues(t, 2, newStateSecond.PlaceEventCount)
}

func TestDoubleBid(t *testing.T) {
	doubleBid := BidPlaced{
		Amount:   12,
		UserId:   4,
		Promoted: true,
	}
	newState := ApplyOnSnapshot(snapshot, doubleBid)

	assert.EqualValues(t, 12, newState.CurrentBid)
	assert.EqualValues(t, 4, newState.CurrentUser)
	assert.EqualValues(t, true, newState.Promoted)
	assert.EqualValues(t, 1, newState.PlaceEventCount)
	assert.EqualValues(t, "Test", newState.Name)

	bidPlaced := BidPlaced{
		Amount:   13,
		UserId:   5,
		Promoted: false,
	}
	newStateSecond := ApplyOnSnapshot(snapshot, bidPlaced)

	assert.EqualValues(t, 13, newStateSecond.CurrentBid)
	assert.EqualValues(t, 5, newStateSecond.CurrentUser)
	assert.EqualValues(t, false, newStateSecond.Promoted)
	assert.EqualValues(t, 1, newStateSecond.PlaceEventCount)
	assert.EqualValues(t, "Test", newStateSecond.Name)
}

func TestBidDeleted(t *testing.T) {
	ss := Auction{
		UUID:            "ttt-ttt",
		Name:            "Test",
		StartDate:       1,
		DueDate:         2,
		CurrentBid:      6,
		CurrentUser:     1,
		PlaceEventCount: 1,
	}

	bidDeleted := BidDeleted{
		ShouldSwap: true,
		Amount:     0,
		UserId:     0,
	}

	newState := ApplyOnSnapshot(ss, bidDeleted)

	assert.EqualValues(t, 0, newState.CurrentBid)
	assert.EqualValues(t, 0, newState.CurrentUser)
	assert.EqualValues(t, 1, newState.DeleteEventCount)
	assert.EqualValues(t, "Test", newState.Name)

	bidDeletedSecond := BidDeleted{
		ShouldSwap: false,
		Amount:     0,
		UserId:     0,
	}

	newStateSecond := ApplyOnSnapshot(newState, bidDeletedSecond)
	assert.EqualValues(t, 0, newStateSecond.CurrentBid)
	assert.EqualValues(t, 0, newStateSecond.CurrentUser)
	assert.EqualValues(t, 2, newStateSecond.DeleteEventCount)
	assert.EqualValues(t, "Test", newStateSecond.Name)
}

func TestNewAuctionFromEvents(t *testing.T) {
	events := []AuctionEvent{
		messageCreate,
		BidPlaced{UserId: 1, Amount: 4, Promoted: false},
		BidPlaced{UserId: 2, Amount: 5, Promoted: false},
		BidPlaced{UserId: 3, Amount: 6, Promoted: false},
		BidPlaced{UserId: 4, Amount: 12, Promoted: true},
		BidPlaced{UserId: 5, Amount: 14, Promoted: false},
		BidDeleted{UserId: 0, Amount: 0, ShouldSwap: false},
		BidDeleted{UserId: 4, Amount: 12, ShouldSwap: true},
		WinnerAnnounced{WinnerId: 4},
	}

	state := NewAuctionFromEvents(events)
	assert.EqualValues(t, 12, state.CurrentBid)
	assert.EqualValues(t, 4, state.CurrentUser)
	assert.EqualValues(t, 2, state.DeleteEventCount)
	assert.EqualValues(t, 5, state.PlaceEventCount)
	assert.EqualValues(t, false, state.Promoted)
	assert.EqualValues(t, 2, state.DeleteEventCount)
	assert.EqualValues(t, 4, state.Winner)
	assert.EqualValues(t, "Test", state.Name)
}
