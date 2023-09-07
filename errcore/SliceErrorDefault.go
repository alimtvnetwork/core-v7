package errcore

import "gitlab.com/auk-go/core/constants"

func SliceErrorDefault(slice *[]string) error {
	return SliceError(constants.NewLineUnix, slice)
}
