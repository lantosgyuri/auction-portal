package comand

import (
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/app"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"time"
)

/* It is currently the same as Auction. I created this abstraction because
if there is a business logic change, than I only should change this struct.
*/
type CreateAuction struct {
	auction domain.CreateAuction
}

type CreateAuctionHandler struct {
	repo app.Repository
}

func (c CreateAuctionHandler) Handle(cmd CreateAuction) error {
	now := int(time.Now().Unix())

	if cmd.auction.Name == "" {
		return errors.New("no name provided for auction")
	}

	if cmd.auction.DueDate < now {
		return errors.New(fmt.Sprintf("not valid DueDate. CurrentTime: %v, DueDate: %v", now, cmd.auction.DueDate))
	}

	if cmd.auction.StartDate < now {
		return errors.New(fmt.Sprintf("not valid StartDate. CurrentTime: %v, StartDate: %v", now, cmd.auction.StartDate))
	}

	if cmd.auction.StartDate > cmd.auction.DueDate {
		return errors.New(fmt.Sprintf("invalid dates. StartDate: %v, DueDate: %v", cmd.auction.StartDate, cmd.auction.DueDate))
	}

	return c.repo.CreateNewAuction(cmd.auction)
}
