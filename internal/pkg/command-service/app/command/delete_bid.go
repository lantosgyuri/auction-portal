package command

import "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"

type DeleteBidHandler struct {
	Repo BidRepository
}

func (d DeleteBidHandler) Handle(bid domain.BidDeleted) error {
	return d.Repo.DeleteBid(bid)
}
