package command

import (
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"time"
)

/* It is currently the same as Auction. I created this abstraction because
if there is a business logic change, than I only should change this struct.
*/
type CreateAuction struct {
	Auction domain.CreateAuction
}

type CreateAuctionHandler struct {
	Repo Repository
}

func (c CreateAuctionHandler) Handle(cmd CreateAuction) error {
	now := int(time.Now().Unix())

	if cmd.Auction.Name == "" {
		return errors.New("no name provided for auction")
	}

	if cmd.Auction.DueDate < now {
		return errors.New(fmt.Sprintf("not valid DueDate. CurrentTime: %v, DueDate: %v", now, cmd.Auction.DueDate))
	}

	if cmd.Auction.StartDate < now {
		return errors.New(fmt.Sprintf("not valid StartDate. CurrentTime: %v, StartDate: %v", now, cmd.Auction.StartDate))
	}

	if cmd.Auction.StartDate > cmd.Auction.DueDate {
		return errors.New(fmt.Sprintf("invalid dates. StartDate: %v, DueDate: %v", cmd.Auction.StartDate, cmd.Auction.DueDate))
	}

	return c.Repo.CreateNewAuction(cmd.Auction)
}
