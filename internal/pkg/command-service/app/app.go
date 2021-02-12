package app

import (
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/command"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateAuction    command.CreateAuctionHandler
	SaveAuctionEvent command.SaveAuctionEventHandler
	UpdateState      command.UpdateStateHandler
	SaveBidEvent     command.SaveBidEventHandler
	CreateBid        command.CreateBidHandler
	DeleteBid        command.DeleteBidHandler
}

type Queries struct {
	GetAuctionState         query.AuctionStateHandler
	GetBidsForAuction       query.BidForAuctionHandler
	GetHighestBidForAuction query.UserHighestBidHandler
}
