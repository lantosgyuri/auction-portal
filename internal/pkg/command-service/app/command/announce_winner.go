package command

import (
	"context"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
)

type AnnounceWinner struct {
	repo StateRepository
}

func (a AnnounceWinner) Handle(ctx context.Context, winnerMessage domain.WinnerAnnounced) error {
	return a.repo.UpdateState(ctx, winnerMessage, func(currentState domain.Auction) domain.Auction {
		return domain.ApplyOnSnapshot(currentState, winnerMessage)
	})
}
