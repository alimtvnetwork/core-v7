package errcore

import "gitlab.com/evatix-go/core/constants"

func SliceErrorDefault(slice *[]string) error {
	return SliceError(constants.NewLineUnix, slice)
}
