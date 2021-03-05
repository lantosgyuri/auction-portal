package command

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type DeleteBid struct {
	Repo BidRepository
}

func (d *DeleteBid) Execute(event domain.Event) error {
	var bid domain.BidDeleted

	if err := marshal.Payload(event, &bid); err != nil {
		return err
	}

	return d.Repo.DeleteBid(bid)
}
