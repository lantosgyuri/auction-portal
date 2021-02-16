package command_test

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

type isHighestAuctionBid = func(ctx context.Context, auctionId string,
	onHighestBid func(topBid domain.Bid, secondBid domain.Bid) error,
) error

type mockedBRepo struct {
	isHighestAuctionBidCallback isHighestAuctionBid
}

func (m mockedBRepo) SaveBidEvent(event domain.BidEventRaw) error {
	return nil
}
func (m mockedBRepo) IsHighestUserBid(context context.Context, bid domain.BidPlaced,
	validate func(userHighestBid domain.Bid) bool,
) bool {
	return false
}
func (m mockedBRepo) IsHighestAuctionBid(ctx context.Context, auctionId string,
	onHighestBid func(topBid domain.Bid, secondBid domain.Bid) error,
) error {
	return m.isHighestAuctionBidCallback(ctx, auctionId, onHighestBid)
}
func (m mockedBRepo) SaveBid(bid domain.Bid) error {
	return nil
}
func (m mockedBRepo) DeleteBid(bid domain.Bid) error {
	return nil
}

type mockedStateRepo struct {
	state domain.Auction
}

func (u *mockedStateRepo) UpdateState(context context.Context, event domain.AuctionEvent,
	update func(auction domain.Auction) (domain.Auction, error),
) error {
	u.state, _ = update(u.state)
	return nil
}

var message = domain.BidDeleted{
	UserId: 4,
	Amount: 5,
}

var mStateRepo = mockedStateRepo{
	state: domain.Auction{},
}

func TestDeleteBidHighestCanFallback(t *testing.T) {
	mockedHighest := func(ctx context.Context, auctionId string,
		onHighestBid func(topBid domain.Bid, secondBid domain.Bid) error,
	) error {
		return onHighestBid(domain.Bid{
			Amount: 5,
			UserId: 4,
		}, domain.Bid{
			Amount: 4,
			UserId: 3,
		})
	}

	handler := createHandler(mockedHighest)

	err := handler.Handle(context.Background(), message)

	assertResults(t, err, mStateRepo.state, 3, 4, 1)
}

func TestDeleteBidHighestCanNotFallback(t *testing.T) {
	mockedHighest := func(ctx context.Context, auctionId string,
		onHighestBid func(topBid domain.Bid, secondBid domain.Bid) error,
	) error {
		return onHighestBid(domain.Bid{
			Amount: 5,
			UserId: 4,
		}, domain.Bid{
			Amount: 0,
			UserId: 0,
		})
	}

	handler := createHandler(mockedHighest)

	err := handler.Handle(context.Background(), message)

	assertResults(t, err, mStateRepo.state, 0, 0, 2)
}

func TestDeleteBidNotHighest(t *testing.T) {
	mockedHighest := func(ctx context.Context, auctionId string,
		onHighestBid func(topBid domain.Bid, secondBid domain.Bid) error,
	) error {
		return onHighestBid(domain.Bid{
			Amount: 8,
			UserId: 9,
		}, domain.Bid{
			Amount: 7,
			UserId: 12,
		})
	}

	handler := createHandler(mockedHighest)

	err := handler.Handle(context.Background(), message)

	assertResults(t, err, mStateRepo.state, 0, 0, 3)
}

func createHandler(callback isHighestAuctionBid) command.DeleteBidHandler {
	mRepo := mockedBRepo{
		isHighestAuctionBidCallback: callback,
	}

	return command.DeleteBidHandler{
		BidRepo:   mRepo,
		StateRepo: &mStateRepo,
	}
}

func assertResults(t *testing.T, err error, state domain.Auction, userId, bidAmount, deletedEventCount int) {
	assert.Nil(t, err)
	assert.EqualValues(t, userId, state.CurrentUser)
	assert.EqualValues(t, bidAmount, mStateRepo.state.CurrentBid)
	assert.EqualValues(t, deletedEventCount, mStateRepo.state.DeleteEventCount)
}
