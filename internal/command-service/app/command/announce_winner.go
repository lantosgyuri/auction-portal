package command

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

type AnnounceWinnerHandler struct {
	Repo StateRepository
}

func (a AnnounceWinnerHandler) Handle(ctx context.Context, message interface{}) error {
	winnerMessage := message.(domain.CreateAuctionRequested)
	return a.Repo.UpdateState(ctx, winnerMessage, func(currentState domain.Auction) (domain.Auction, error) {
		return domain.ApplyOnSnapshot(currentState, winnerMessage), nil
	})
}
