package command_service

type Publisher interface {
	Publish() error
}

type EventDBRepository interface {
	SaveEvent() error
	SaveSnapshot() error
}

type EventCreator struct {
	publisher Publisher
}

func SaveToStore(event interface{}) {}

func SaveSnapshot(snapshot interface{}) {}

func Publish(event interface{}) {
	// Create new event
}
