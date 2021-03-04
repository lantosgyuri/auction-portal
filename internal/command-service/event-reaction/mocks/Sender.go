// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "github.com/lantosgyuri/auction-portal/internal/command-service/domain"

import mock "github.com/stretchr/testify/mock"

// Sender is an autogenerated mock type for the Sender type
type Sender struct {
	mock.Mock
}

// NotifyUserFail provides a mock function with given fields: notifyEvent
func (_m *Sender) NotifyUserFail(notifyEvent domain.NotifyEvent) {
	_m.Called(notifyEvent)
}

// NotifyUserSuccess provides a mock function with given fields: notifyEvent
func (_m *Sender) NotifyUserSuccess(notifyEvent domain.NotifyEvent) {
	_m.Called(notifyEvent)
}

// PublishData provides a mock function with given fields: event
func (_m *Sender) PublishData(event domain.Event) {
	_m.Called(event)
}
