package command

import "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"

type SaveWinner struct {
	Repo AuctionRepository
}

func (s SaveWinner) Handle(winner domain.AuctionWinnerMessage) error {
	return s.Repo.SaveWinner(winner)
}
