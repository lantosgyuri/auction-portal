package command_service

type AggregateValidator interface {
	Validate() error
}

type EventCreator struct{}
