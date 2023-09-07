package errcore

import (
	"errors"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func SliceToErrorPtr(errorSlice *[]string) error {
	if errorSlice == nil || len(*errorSlice) == 0 {
		return nil
	}

	fullError := strings.Join(
		*errorSlice,
		constants.NewLineUnix)

	return errors.New(fullError)
}
