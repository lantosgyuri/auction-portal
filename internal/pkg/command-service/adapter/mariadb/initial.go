package mariadb

import (
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"sync"
)

func MigrateSotDb(wg *sync.WaitGroup) {

	if connection.SotDb == nil {
		fmt.Print("No DB is set up")
		wg.Done()
	}

}
