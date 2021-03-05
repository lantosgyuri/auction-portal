package command

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/marshal"
)

type SaveWinner struct {
	Repo AuctionRepository
}

func (s *SaveWinner) Execute(event domain.Event) error {
	var winner domain.WinnerAnnounced

	if err := marshal.Payload(event, &winner); err != nil {
		return err
	}

	return s.Repo.SaveWinner(winner)
}
