package command_service

type ValidatorDBRepository interface {
	GetSnapshot() (interface{}, error)
}

type Validator struct {
	repository ValidatorDBRepository
}

func Validate(event interface{}) (interface{}, error) {
	// Get Snapshot
	// Validate event
	// Return New Snapshot or error
	return nil, nil
}
