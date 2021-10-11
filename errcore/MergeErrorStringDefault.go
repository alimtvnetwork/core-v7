package errcore

import "gitlab.com/evatix-go/core/constants"

func MergeErrorStringDefault(
	errorItems ...error,
) string {
	if errorItems == nil {
		return ""
	}

	return MergeErrorString(constants.Space, errorItems...)
}
