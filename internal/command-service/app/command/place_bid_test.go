package command_test

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

type isHighestBid = func(context context.Context, bid domain.BidPlaced,
	validate func(userHighestBid domain.Bid) bool,
) bool

type updateState = func(context context.Context, event domain.AuctionEvent,
	update func(auction domain.Auction) (domain.Auction, error),
) error

type mockedBidRepo struct {
	highestBidCallback  isHighestBid
	updateStateCallback updateState
}

func (m mockedBidRepo) SaveBidEvent(event domain.BidEventRaw) error {
	return nil
}

func (m mockedBidRepo) IsHighestUserBid(context context.Context, bid domain.BidPlaced,
	validate func(userHighestBid domain.Bid) bool,
) bool {
	return m.highestBidCallback(context, bid, validate)
}

func (m mockedBidRepo) IsHighestAuctionBid(ctx context.Context, auctionId string,
	onHighestBid func(topBids []domain.Bid) error,
) error {
	return nil
}

func (m mockedBidRepo) SaveBid(bid domain.Bid) error {
	return nil
}
func (m mockedBidRepo) DeleteBid(bid domain.Bid) error {
	return nil
}

func (m mockedBidRepo) UpdateState(context context.Context, event domain.AuctionEvent,
	update func(auction domain.Auction) (domain.Auction, error)) error {
	return m.updateStateCallback(context, event, update)
}

func TestBidSmallerThanBefore(t *testing.T) {
	mockedBidFunction := func(context context.Context, bid domain.BidPlaced,
		validate func(userHighestBid domain.Bid) bool,
	) bool {
		return validate(domain.Bid{Amount: 1000})
	}

	handler := command.PlaceBidHandler{
		BidRepo: mockedBidRepo{
			highestBidCallback: mockedBidFunction,
		},
		StateRepo: mockedBidRepo{},
	}

	message := domain.BidPlaced{
		Amount: 1,
	}

	err := handler.Handle(context.Background(), message)

	assert.NotNil(t, err)
	assert.EqualValues(t, "bid is smaller than latest bid", err.Error())
}

func TestBidIsSmallerThanCurrentBid(t *testing.T) {
	mockedBidFunction := func(context context.Context, bid domain.BidPlaced,
		validate func(userHighestBid domain.Bid) bool,
	) bool {
		return true
	}

	mockedUpdateFunction := func(context context.Context, event domain.AuctionEvent,
		update func(auction domain.Auction) (domain.Auction, error)) error {
		_, err := update(domain.Auction{
			CurrentBid: 1000,
		})
		return err
	}

	handler := command.PlaceBidHandler{
		BidRepo: mockedBidRepo{
			highestBidCallback: mockedBidFunction,
		},
		StateRepo: mockedBidRepo{
			updateStateCallback: mockedUpdateFunction,
		},
	}

	message := domain.BidPlaced{
		Amount: 1,
	}

	err := handler.Handle(context.Background(), message)

	assert.NotNil(t, err)
	assert.EqualValues(t, "ignore bid. Bid is smaller or equal with current bid", err.Error())
}
