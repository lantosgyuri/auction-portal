package command

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type SaveBid struct {
	Repo BidRepository
}

func (s *SaveBid) Execute(event domain.Event) error {
	var bid domain.BidPlaced

	if err := marshal.Payload(event, &bid); err != nil {
		return err
	}

	return s.Repo.SaveBid(bid)
}
