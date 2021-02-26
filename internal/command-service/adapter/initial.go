package adapter

import (
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"sync"
)

func MigrateSotDb(wg *sync.WaitGroup) {

	db := connection.GetMariDbConnection()

	if err := db.AutoMigrate(
		&domain.AuctionEventRaw{},
		&domain.BidEventRaw{},
		&domain.UserEventRaw{},
		&domain.Auction{},
		&domain.User{},
		&domain.Bid{},
	); err != nil {
		fmt.Print("can not migrate SotDb")
		wg.Done()
	}
}
