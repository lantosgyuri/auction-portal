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
	Commands[domain.BidDeleteRequested] = BidDeleteRequestedCommand{}
}

type BidDeleteRequestedCommand struct{}

func (b BidDeleteRequestedCommand) Execute(application app.Application, event domain.Event) error {
	var bidDeleteMessage domain.BidDeleted

	if err := json.Unmarshal(event.Payload, &bidDeleteMessage); err != nil {
		return errors.New(fmt.Sprintf("Error happened with unmarshalling winner message: %v", err))
	}

	if err := UpdateState(context.Background(), bidDeleteMessage, application); err != nil {
		return err
	}
	if err := application.Commands.SaveBidEvent.Handle(event.Event, bidDeleteMessage); err != nil {
		return err
	}

	return nil
}
