package domain

type Event struct {
	Event   string
	Payload []byte
}

type NormalizedAuctionEvent struct {
	Event string
	Data  string
}
