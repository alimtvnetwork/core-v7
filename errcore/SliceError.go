package errcore

import (
	"errors"
	"strings"
)

func SliceError(sep string, slice *[]string) error {
	if slice == nil || len(*slice) == 0 {
		return nil
	}

	msg := strings.Join(*slice, sep)

	return errors.New(msg)
}
