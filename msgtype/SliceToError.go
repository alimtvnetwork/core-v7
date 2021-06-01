package msgtype

import (
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func SliceToError(errorSlice *[]string) error {
	if errorSlice == nil || len(*errorSlice) == 0 {
		return nil
	}

	fullError := strings.Join(
		*errorSlice,
		constants.NewLineUnix)

	return errors.New(fullError)
}
