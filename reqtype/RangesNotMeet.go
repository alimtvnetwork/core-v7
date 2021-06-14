package reqtype

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/msgtype"
)

func RangesNotMeet(
	message string,
	reqs ...Request,
) string {
	if len(reqs) == 0 {
		return constants.EmptyString
	}

	currentStart := start(&reqs)
	currentEnd := end(&reqs)

	return msgtype.RangeNotMeet(
		message,
		currentStart,
		currentEnd,
		reqs)
}
