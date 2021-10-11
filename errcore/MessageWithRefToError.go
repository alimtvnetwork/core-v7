package errcore

import (
	"errors"
	"fmt"
)

func MessageWithRefToError(msg string, reference interface{}) error {
	return errors.New(fmt.Sprintf(
		messageMapFormat,
		msg,
		reference),
	)
}
