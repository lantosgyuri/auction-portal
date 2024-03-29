// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "github.com/lantosgyuri/auction-portal/internal/command-service/domain"

import mock "github.com/stretchr/testify/mock"

// PreserveBidEvent is an autogenerated mock type for the PreserveBidEvent type
type PreserveBidEvent struct {
	mock.Mock
}

// Handle provides a mock function with given fields: eventName, event
func (_m *PreserveBidEvent) Handle(eventName string, event domain.BidEvent) error {
	ret := _m.Called(eventName, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.BidEvent) error); ok {
		r0 = rf(eventName, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
