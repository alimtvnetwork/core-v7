package msgtype

import "fmt"

func EnumRangeNotMeet(
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

	return EnumValuesOutOfRange.String() +
		ReferenceStart +
		rangeStr +
		ReferenceEnd
}
