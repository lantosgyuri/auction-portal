package command

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type AnnounceWinnerHandler struct {
	Repo StateRepository
}

func (a AnnounceWinnerHandler) Handle(ctx context.Context, winnerMessage domain.WinnerAnnounced) error {
	return a.Repo.UpdateState(ctx, winnerMessage, func(currentState domain.Auction) (domain.Auction, error) {
		return domain.ApplyOnSnapshot(currentState, winnerMessage), nil
	})
}
