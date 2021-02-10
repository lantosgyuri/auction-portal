package event_reaction

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

func init() {
	Commands[EventMessageNames.AuctionRequested] = AuctionRequested{}
}

type AuctionRequested struct{}

func (a AuctionRequested) Execute(application app.Application, event domain.Event) error {
	var auction domain.CreateAuction
	if err := json.Unmarshal(event.Payload, &auction); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling user: %v", err))
	}
	err := application.Commands.CreateAuction.Handle(command.CreateAuction{Auction: auction})
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened with creating auction: %v", err))
	}
	err = application.Commands.SaveAuctionEvent.Handle(event)
	if err != nil {
		return errors.New(fmt.Sprintf("Error happened during saving the auction event: %v", err))
	}
	fmt.Printf("Event is: %v", auction)

	return nil
}
