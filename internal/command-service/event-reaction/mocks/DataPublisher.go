// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "github.com/lantosgyuri/auction-portal/internal/command-service/domain"

import mock "github.com/stretchr/testify/mock"

// DataPublisher is an autogenerated mock type for the DataPublisher type
type DataPublisher struct {
	mock.Mock
}

// PublishData provides a mock function with given fields: event
func (_m *DataPublisher) PublishData(event domain.Event) {
	_m.Called(event)
}
