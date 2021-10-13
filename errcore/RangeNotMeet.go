package errcore

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func RangeNotMeet(
	otherMsg string,
	rangeStart interface{},
	rangeEnd interface{},
	wholeRange interface{},
) string {
	rangeStr := ""

	if wholeRange == nil {
		rangeStr = fmt.Sprintf(rangeWithoutRangeFormat, rangeStart, rangeEnd)
	} else {
		rangeStr = fmt.Sprintf(rangeWithRangeFormat, rangeStart, rangeEnd, wholeRange)
	}

	return OutOfRangeType.String() +
		constants.Space +
		otherMsg +
		ReferenceStart +
		rangeStr +
		ReferenceEnd
}
