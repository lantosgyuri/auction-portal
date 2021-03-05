package custom_error

import (
	"errors"
	"fmt"
)

func Create(text string, err error) error {
	return errors.New(fmt.Sprintf("can not update auction: %v", err))
}
