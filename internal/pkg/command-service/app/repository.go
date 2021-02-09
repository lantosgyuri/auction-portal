package app

import "github.com/lantosgyuri/auction-portal/internal/pkg/command-service/domain"

type Repository interface {
	SaveAuctionEvent(event domain.NormalizedAuctionEvent) error
	SaveBidEvent() error
	SaveUserEvent() error
	CreateNewAuction(event domain.CreateAuction) error
	CreateNewUser() error
}
