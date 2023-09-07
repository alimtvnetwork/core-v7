package errcore

import (
	"gitlab.com/auk-go/core/constants"
)

func ToString(err error) string {
	if err == nil {
		return constants.EmptyString
	}

	return err.Error()
}
