package csvinternal

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// RangeNamesWithValuesIndexes
//
//  Returns a new slice where
//  format
//   - `name[ValueIndex]`
//  example
//   - `SomeName[1]`
func RangeNamesWithValuesIndexes(
	rangedItems ...string,
) []string {
	if len(rangedItems) == 0 {
		return []string{}
	}

	compiledRanges := make([]string, len(rangedItems))

	for i, item := range rangedItems {
		compiledRanges[i] = fmt.Sprintf(
			constants.EnumNameValueFormat,
			item,
			i)
	}

	return compiledRanges
}
