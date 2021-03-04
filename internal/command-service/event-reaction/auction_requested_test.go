package event_reaction_test

import (
	"encoding/json"
	"errors"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	event_reaction "github.com/lantosgyuri/auction-portal/internal/command-service/event-reaction"
	"github.com/lantosgyuri/auction-portal/internal/command-service/event-reaction/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type AuctionRequestedSuite struct {
	suite.Suite

	handler   *mocks.AuctionCreateEventHandler
	preserver *mocks.AuctionEventPreserver
	sender    *mocks.Sender

	auction domain.CreateAuctionRequested
	command event_reaction.AuctionRequestedCommand
	event   domain.Event
}

func (a *AuctionRequestedSuite) SetupTest() {
	a.handler = new(mocks.AuctionCreateEventHandler)
	a.preserver = new(mocks.AuctionEventPreserver)
	a.sender = new(mocks.Sender)

	a.auction = domain.CreateAuctionRequested{
		DueDate:   int(time.Now().AddDate(0, 0, 8).Unix()),
		StartDate: int(time.Now().AddDate(0, 0, 1).Unix()),
		Name:      "Test6",
	}

	auctionBytes, _ := json.Marshal(a.auction)

	a.event = domain.Event{
		Event:         domain.AuctionRequested,
		CorrelationId: 1234,
		Payload:       auctionBytes,
	}
	a.command = event_reaction.CreateAuctionRequestCommandWithInterfaces(a.handler, a.preserver, a.sender)
}

func (a *AuctionRequestedSuite) TestPreserveFailed() {
	err := errors.New("error")
	notifyEvent := domain.NotifyEvent{
		CorrelationId: 1234,
		Event:         "Create_auction",
		Success:       false,
		Error:         "error happened with saving data: error",
	}

	a.preserver.On("Handle", domain.AuctionRequested, a.auction).Return(err)
	a.sender.On("NotifyUserFail", notifyEvent)

	a.command.Execute(a.event)
	a.sender.AssertNumberOfCalls(a.T(), "NotifyUserFail", 1)
	a.sender.AssertNumberOfCalls(a.T(), "NotifyUserSuccess", 0)
	a.sender.AssertNumberOfCalls(a.T(), "PublishData", 0)

	a.sender.AssertExpectations(a.T())
	a.preserver.AssertExpectations(a.T())
}

func (a *AuctionRequestedSuite) TestHandleFailed() {
	err := errors.New("error")
	notifyEvent := domain.NotifyEvent{
		CorrelationId: 1234,
		Event:         "Create_auction",
		Success:       false,
		Error:         "error happened with auction creating: error",
	}

	a.preserver.On("Handle", domain.AuctionRequested, a.auction).Return(nil)
	a.handler.On("Handle", a.auction).Return(err)
	a.sender.On("NotifyUserFail", notifyEvent)

	a.command.Execute(a.event)
	a.sender.AssertNumberOfCalls(a.T(), "NotifyUserFail", 1)
	a.sender.AssertNumberOfCalls(a.T(), "NotifyUserSuccess", 0)
	a.sender.AssertNumberOfCalls(a.T(), "PublishData", 0)

	a.sender.AssertExpectations(a.T())
	a.preserver.AssertExpectations(a.T())
	a.handler.AssertExpectations(a.T())
}

func (a *AuctionRequestedSuite) TestSuccess() {
	notifyEvent := domain.NotifyEvent{
		CorrelationId: 1234,
		Event:         "Create_auction",
		Success:       true,
	}

	a.preserver.On("Handle", domain.AuctionRequested, a.auction).Return(nil)
	a.handler.On("Handle", a.auction).Return(nil)
	a.sender.On("NotifyUserSuccess", notifyEvent)
	a.sender.On("PublishData", a.event)

	a.command.Execute(a.event)
	a.sender.AssertNumberOfCalls(a.T(), "NotifyUserFail", 0)
	a.sender.AssertNumberOfCalls(a.T(), "NotifyUserSuccess", 1)
	a.sender.AssertNumberOfCalls(a.T(), "PublishData", 1)

	a.sender.AssertExpectations(a.T())
	a.preserver.AssertExpectations(a.T())
	a.handler.AssertExpectations(a.T())
}

func TestCreateAuctionRequestedCommand(t *testing.T) {
	suite.Run(t, new(AuctionRequestedSuite))
}
