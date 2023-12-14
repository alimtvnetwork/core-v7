package corecsv

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// AnyToTypesCsvDefault
//
// if references empty or len 0 then empty string returned.
//
// Formats :
//   - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//   - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//   - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func AnyToTypesCsvDefault(
	references ...interface{},
) string {
	toSlice := AnyToTypesCsvStrings(
		false,
		false,
		references...)

	return strings.Join(
		toSlice,
		constants.CommaSpace,
	)
}
