package command

import "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"

type AuctionRepository interface {
	SaveAuctionEvent(event domain.AuctionEvent) error
	CreateNewAuction(auction domain.Auction) error
	SaveWinner(auctionWinnerMessage domain.WinnerAnnounced) error
}
