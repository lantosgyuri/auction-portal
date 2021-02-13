package command

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type AnnounceWinner struct {
	Repo StateRepository
}

func (a AnnounceWinner) Handle(ctx context.Context, winnerMessage domain.WinnerAnnounced) error {
	return a.Repo.UpdateState(ctx, winnerMessage, func(currentState domain.Auction) (domain.Auction, error) {
		return domain.ApplyOnSnapshot(currentState, winnerMessage), nil
	})
}
