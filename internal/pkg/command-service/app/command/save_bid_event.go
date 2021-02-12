package command

import (
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type SaveBidEventHandler struct {
	Repo BidRepository
}

func (s SaveBidEventHandler) Handle(event domain.Event) error {
	return nil
}
