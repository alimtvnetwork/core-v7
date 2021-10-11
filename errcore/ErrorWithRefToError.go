package errcore

import (
	"errors"
)

func ErrorWithRefToError(err error, reference interface{}) error {
	if err == nil {
		return nil
	}

	return errors.New(ErrorWithRef(err, reference))
}
