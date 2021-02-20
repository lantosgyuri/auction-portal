package command_test

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command/mocks"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var handler = command.CreateAuctionHandler{
	Repo: &mocks.AuctionRepository{},
}

type errorTest struct {
	message      domain.CreateAuctionRequested
	errorMessage string
}

func TestErrorCases(t *testing.T) {

	errorCases := map[string]errorTest{
		"NoNameProvided": {
			message: domain.CreateAuctionRequested{
				Name: "",
			},
			errorMessage: "no name provided for auction",
		},

		"DuDateBeforeNow": {
			message: domain.CreateAuctionRequested{
				DueDate:   int(time.Now().AddDate(0, 0, -1).Unix()),
				StartDate: int(time.Now().AddDate(0, 0, 2).Unix()),
				Name:      "Test",
			},
			errorMessage: "not valid DueDate.",
		},
		"StartDateBeforeNow": {
			message: domain.CreateAuctionRequested{
				DueDate:   int(time.Now().AddDate(0, 0, 7).Unix()),
				StartDate: int(time.Now().AddDate(0, 0, -1).Unix()),
				Name:      "Test",
			},
			errorMessage: "not valid StartDate.",
		},
		"StartDateBeforeDueDate": {
			message: domain.CreateAuctionRequested{
				DueDate:   int(time.Now().AddDate(0, 0, 6).Unix()),
				StartDate: int(time.Now().AddDate(0, 0, 7).Unix()),
				Name:      "Test",
			},
			errorMessage: "invalid dates",
		},
	}

	for _, value := range errorCases {
		err := handler.Handle(value.message)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), value.errorMessage)
	}

}
