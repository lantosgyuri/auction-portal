package command

import "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"

type UpdateStateCommand struct {
	CurrentState domain.Auction
	Event        domain.AuctionEvent
}

type UpdateState struct {
	Repo AuctionRepository
}

func (u UpdateState) Handle(cmd UpdateStateCommand) error {
	newState := domain.ApplyOnSnapshot(cmd.CurrentState, cmd.Event)
	return u.Repo.UpdateAuctionState(newState)
}
