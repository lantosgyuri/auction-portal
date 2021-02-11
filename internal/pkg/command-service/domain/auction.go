package domain

type AuctionEvent interface {
	isAuctionEvent()
}

type CreateAuctionRequested struct {
	Name      string
	DueDate   int
	StartDate int
	Timestamp int
}

func (c CreateAuctionRequested) isAuctionEvent() {}

type WinnerAnnounced struct {
	AuctionId int
	WinnerId  int
	Timestamp int
}

func (a WinnerAnnounced) isAuctionEvent() {}

type Auction struct {
	Id               int
	Name             string
	DueDate          int
	StartDate        int
	Version          int
	Winner           int
	CurrentBid       int
	PlaceEventCount  int
	DeleteEventCount int
}

func NewAuctionFromEvents(events []AuctionEvent) Auction {
	a := Auction{}

	for _, event := range events {
		a.On(event)
	}

	return a
}

func New(message CreateAuctionRequested) Auction {
	a := Auction{}
	a.On(message)
	return a
}

func (a *Auction) On(event AuctionEvent) {
	switch e := event.(type) {
	case *CreateAuctionRequested:
		a.Name = e.Name
		a.DueDate = e.DueDate
		a.StartDate = e.StartDate
	case *WinnerAnnounced:

	case *BidDeleted:
	case *BidDoubled:
	case *BidPlaced:
	}

	a.Version++
}
