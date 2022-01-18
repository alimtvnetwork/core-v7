package errcore

import "gitlab.com/evatix-go/core/constants"

func SimpleHandleErrMany(msg string, errorItems ...error) {
	if errorItems == nil || len(errorItems) == 0 {
		return
	}

	singleErr := ManyErrorToSingle(errorItems)

	if singleErr != nil {
		panic(singleErr.Error() + constants.NewLineUnix + msg)
	}
}
