package marshal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
)

func Payload(event domain.Event, target interface{}) error {
	err := json.Unmarshal(event.Payload, target)

	if err != nil {
		return errors.New(fmt.Sprintf("can not unmarshal payload: %v", err))
	}

	return nil
}
