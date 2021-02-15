package mariadb

import (
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"sync"
)

func MigrateSotDb(wg *sync.WaitGroup) {

	if connection.SotDb == nil {
		fmt.Print("No DB is set up")
		wg.Done()
	}

	if err := connection.SotDb.AutoMigrate(
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
