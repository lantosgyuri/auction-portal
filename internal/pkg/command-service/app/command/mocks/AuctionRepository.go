// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
import mock "github.com/stretchr/testify/mock"

// AuctionRepository is an autogenerated mock type for the AuctionRepository type
type AuctionRepository struct {
	mock.Mock
}

// CreateNewAuction provides a mock function with given fields: auction
func (_m *AuctionRepository) CreateNewAuction(auction domain.Auction) error {
	ret := _m.Called(auction)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Auction) error); ok {
		r0 = rf(auction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveAuctionEvent provides a mock function with given fields: event
func (_m *AuctionRepository) SaveAuctionEvent(event domain.AuctionEventRaw) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.AuctionEventRaw) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
