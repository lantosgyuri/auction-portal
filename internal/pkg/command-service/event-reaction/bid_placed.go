package event_reaction

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

func init() {
	Commands[domain.BidPlaceRequested] = BidPlaceRequestedCommand{}
}

type BidPlaceRequestedCommand struct{}

func (b BidPlaceRequestedCommand) Execute(application app.Application, event domain.Event) error {
	var bidPlacedMessage domain.BidPlaced
	if err := json.Unmarshal(event.Payload, &bidPlacedMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling winner message: %v", err))
	}

	if err := UpdateState(context.Background(), bidPlacedMessage, application); err != nil {
		return err
	}

	if err := application.Commands.CreateBid.Handle(bidPlacedMessage); err != nil {
		return err
	}
	if err := application.Commands.SaveBidEvent.Handle(event.Event, bidPlacedMessage); err != nil {
		return err
	}

	return nil
}
