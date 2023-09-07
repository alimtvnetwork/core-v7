package errcore

import "gitlab.com/auk-go/core/constants"

func SimpleHandleErr(err error, msg string) {
	if err == nil {
		return
	}

	panic(err.Error() + constants.NewLineUnix + msg)
}
