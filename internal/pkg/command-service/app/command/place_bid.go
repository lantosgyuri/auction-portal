package command

import "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"

type CreateBidHandler struct {
	Repo BidRepository
}

func (c CreateBidHandler) Handle(bid domain.BidPlaced) error {
	return c.Repo.CreateBid(bid)
}
