package errcore

import (
	"gitlab.com/auk-go/core/constants"
)

func MergeErrorsToStringDefault(
	errorItems ...error,
) string {
	if errorItems == nil {
		return ""
	}

	return MergeErrorsToString(constants.Space, errorItems...)
}
