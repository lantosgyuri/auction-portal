package domain

type AuctionEvent struct {
	EventType string
	Name      string
	AuctionId int
	DueDate   int
	StartDate int
	Winner    int
	Timestamp int
}

type CreateAuctionMessage struct {
	Name      string
	DueDate   int
	StartDate int
	Timestamp int
}

type AuctionWinnerMessage struct {
	AuctionId int
	WinnerId  int
	Timestamp int
}
