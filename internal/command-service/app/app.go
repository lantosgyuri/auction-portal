package app

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/app/command"
)

type Application struct {
	Commands Commands
}

type Commands struct {
	CreateAuction    command.CreateAuctionHandler
	SaveAuctionEvent command.SaveAuctionEventHandler
	SaveBidEvent     command.SaveBidEventHandler
	SaveUserEvent    command.SaveUserEventHandler
	PlaceBid         command.PlaceBidHandler
	DeleteBid        command.DeleteBidHandler
	AnnounceWinner   command.AnnounceWinnerHandler
	CreateUser       command.CreateUserHandler
	DeleteUser       command.DeleteUserHandler
}
