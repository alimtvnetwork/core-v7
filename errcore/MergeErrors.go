package errcore

import (
	"gitlab.com/evatix-go/core/constants"
)

func MergeErrors(errorItems ...error) error {
	if len(errorItems) == 0 {
		return nil
	}

	sliceErr := make(
		[]string,
		constants.Zero,
		len(errorItems))

	for _, err := range errorItems {
		if err == nil {
			continue
		}

		sliceErr = append(sliceErr, err.Error())
	}

	return SliceToError(sliceErr)
}
