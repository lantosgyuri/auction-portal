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
	SaveWinner       command.SaveWinner
}

type Queries struct {
	GetBidsForAuction query.BidsForAuctionHandler
}
