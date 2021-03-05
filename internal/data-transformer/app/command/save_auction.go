package command

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type AuctionPreserver struct {
	Repo AuctionRepository
}

func (a *AuctionPreserver) Execute(event domain.Event) error {
	var auction domain.CreateAuctionRequested

	if err := marshal.Payload(event, &auction); err != nil {
		return err
	}

	return a.Repo.SaveAuction(auction)
}
