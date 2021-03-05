package marshal

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Payload(payload []byte, target interface{}) error {
	err := json.Unmarshal(payload, target)

	if err != nil {
		return errors.New(fmt.Sprintf("can not unmarshal payload: %v", err))
	}

	return nil
}
