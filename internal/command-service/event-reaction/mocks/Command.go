// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "github.com/lantosgyuri/auction-portal/internal/command-service/domain"

import mock "github.com/stretchr/testify/mock"

// Command is an autogenerated mock type for the Command type
type Command struct {
	mock.Mock
}

// Execute provides a mock function with given fields: event
func (_m *Command) Execute(event domain.Event) {
	_m.Called(event)
}
