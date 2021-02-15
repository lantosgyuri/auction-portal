package command

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"
	"time"
)

type CreateAuctionHandler struct {
	Repo AuctionRepository
}

func (c CreateAuctionHandler) Handle(auction domain.CreateAuctionRequested) error {
	now := int(time.Now().Unix())

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

	auction.UUID = uuid.New().String()
	newAuction := domain.NewAuction(auction)

	return c.Repo.CreateNewAuction(newAuction)
}
