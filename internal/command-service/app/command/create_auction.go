package command

import (
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"time"
)

type CreateAuctionHandler struct {
	Repo AuctionRepository
}

func (c CreateAuctionHandler) Handle(auction domain.CreateAuctionRequested) error {
	now := int(time.Now().Unix())

	// Create connection

	if auction.Name == "" {
		return errors.New("no name provided for auction")
	}

	if auction.DueDate < now {
		return errors.New(fmt.Sprintf("not valid DueDate. CurrentTime: %v, DueDate: %v", now, auction.DueDate))
	}

	if auction.StartDate < now {
		return errors.New(fmt.Sprintf("not valid StartDate. CurrentTime: %v, StartDate: %v", now, auction.StartDate))
	}

	if auction.StartDate > auction.DueDate {
		return errors.New(fmt.Sprintf("invalid dates. StartDate: %v, DueDate: %v", auction.StartDate, auction.DueDate))
	}

	newAuction := domain.NewAuction(auction)

	return c.Repo.CreateNewAuction(newAuction)
}
