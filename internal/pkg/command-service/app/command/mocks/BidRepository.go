// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import domain "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
import mock "github.com/stretchr/testify/mock"

// BidRepository is an autogenerated mock type for the BidRepository type
type BidRepository struct {
	mock.Mock
}

// DeleteBid provides a mock function with given fields: bid
func (_m *BidRepository) DeleteBid(bid domain.BidDeleted) error {
	ret := _m.Called(bid)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.BidDeleted) error); ok {
		r0 = rf(bid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsHighestAuctionBid provides a mock function with given fields: ctx, auctionId, onHighestBid
func (_m *BidRepository) IsHighestAuctionBid(ctx context.Context, auctionId string, onHighestBid func(domain.Bid, domain.Bid) error) error {
	ret := _m.Called(ctx, auctionId, onHighestBid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, func(domain.Bid, domain.Bid) error) error); ok {
		r0 = rf(ctx, auctionId, onHighestBid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsHighestUserBid provides a mock function with given fields: _a0, bid, validate
func (_m *BidRepository) IsHighestUserBid(_a0 context.Context, bid domain.BidPlaced, validate func(domain.Bid) bool) bool {
	ret := _m.Called(_a0, bid, validate)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, domain.BidPlaced, func(domain.Bid) bool) bool); ok {
		r0 = rf(_a0, bid, validate)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SaveBid provides a mock function with given fields: bid
func (_m *BidRepository) SaveBid(bid domain.BidPlaced) error {
	ret := _m.Called(bid)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.BidPlaced) error); ok {
		r0 = rf(bid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveBidEvent provides a mock function with given fields: event
func (_m *BidRepository) SaveBidEvent(event domain.BidEventRaw) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.BidEventRaw) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
