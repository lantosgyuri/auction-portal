package app

import (
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/command"
)

type Application struct {
	Commands Commands
}

type Commands struct {
	CreateAuction    command.CreateAuctionHandler
	SaveAuctionEvent command.SaveAuctionEventHandler
	SaveBidEvent     command.SaveBidEventHandler
	PlaceBid         command.PlaceBidHandler
	DeleteBid        command.DeleteBidHandler
	AnnounceWinner   command.AnnounceWinner
}
