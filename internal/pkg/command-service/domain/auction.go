package domain

type AuctionEvent interface {
	GetAuctionId() string
}

type AuctionEventRaw struct {
	EventType string
	Name      string
	AuctionId string
	DueDate   int
	StartDate int
	Winner    int
	Timestamp int
}

type CreateAuctionRequested struct {
	UUID      string
	Name      string
	DueDate   int
	StartDate int
	Timestamp int
}

func (c CreateAuctionRequested) GetAuctionId() string {
	return c.UUID
}

type WinnerAnnounced struct {
	AuctionId string
	WinnerId  int
	Timestamp int
}

func (w WinnerAnnounced) GetAuctionId() string {
	return w.AuctionId
}

type Auction struct {
	UUID             string
	Name             string
	DueDate          int
	StartDate        int
	Winner           int
	CurrentBid       int
	CurrentUser      int
	Promoted         bool
	PlaceEventCount  int
	DeleteEventCount int
}

func NewAuctionFromEvents(events []AuctionEvent) Auction {
	a := Auction{}

	for _, event := range events {
		a.Apply(event)
	}

	return a
}

func ApplyOnSnapshot(snapshot Auction, event AuctionEvent) Auction {
	snapshot.Apply(event)
	return snapshot
}

func NewAuction(message CreateAuctionRequested) Auction {
	a := Auction{}
	a.Apply(message)
	return a
}

func (a *Auction) Apply(event AuctionEvent) {
	switch e := event.(type) {
	case CreateAuctionRequested:
		a.UUID = e.UUID
		a.Name = e.Name
		a.DueDate = e.DueDate
		a.StartDate = e.StartDate
	case WinnerAnnounced:
		a.Winner = e.WinnerId
	case BidDeleted:
		if e.ShouldSwap {
			a.CurrentUser = e.UserId
			a.CurrentBid = e.Amount
		}
		a.DeleteEventCount++
	case BidPlaced:
		if e.Promoted {
			a.Promoted = true
			a.CurrentUser = e.UserId
			a.CurrentBid = e.Amount
			a.PlaceEventCount++
			break
		}
		a.Promoted = false
		a.CurrentUser = e.UserId
		a.CurrentBid = e.Amount
		a.PlaceEventCount++
	}

}
