package errcore

import "gitlab.com/evatix-go/core/constants"

func SimpleHandleErr(err error, msg string) {
	if err == nil {
		return
	}

	panic(err.Error() + constants.NewLineUnix + msg)
}

func ManyErrorToSingleDirect(errorItems ...error) error {
	if errorItems == nil || len(errorItems) == 0 {
		return nil
	}

	return ManyErrorToSingle(errorItems)
}

func ManyErrorToSingle(errorItems []error) error {
	if errorItems == nil || len(errorItems) == 0 {
		return nil
	}

	sliceErr := make([]string, 0, len(errorItems))

	for _, err := range errorItems {
		if err == nil {
			continue
		}

		sliceErr = append(sliceErr, err.Error())
	}

	return SliceToError(sliceErr)
}

func SimpleHandleErrMany(msg string, errorItems ...error) {
	if errorItems == nil || len(errorItems) == 0 {
		return
	}

	singleErr := ManyErrorToSingle(errorItems)

	if singleErr != nil {
		panic(singleErr.Error() + constants.NewLineUnix + msg)
	}
}
