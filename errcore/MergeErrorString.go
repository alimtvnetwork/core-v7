package errcore

import "strings"

func MergeErrorString(
	joiner string,
	errorItems ...error,
) string {
	if errorItems == nil {
		return ""
	}

	slice := SliceErrorsToStrings(errorItems...)

	return strings.Join(slice, joiner)
}
