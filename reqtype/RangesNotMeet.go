package reqtype

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

func RangesNotMeet(
	message string,
	reqs ...Request,
) string {
	if len(reqs) == 0 {
		return constants.EmptyString
	}

	currentStart := start(reqs)
	currentEnd := end(reqs)

	return errcore.RangeNotMeet(
		message,
		currentStart,
		currentEnd,
		reqs)
}
