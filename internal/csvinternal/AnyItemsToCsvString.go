package csvinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// AnyItemsToCsvString
//
// if references empty or len 0 then empty string returned.
//
// Final join whole lines with the joiner given (... joiner item)
//
// Formats :
//  - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//  - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//  - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func AnyItemsToCsvString(
	joiner string,
	isIncludeQuote,
	isIncludeSingleQuote bool,
	references ...interface{},
) string {
	if len(references) == 0 {
		return constants.EmptyString
	}

	slice := AnyItemsToCsvStrings(
		isIncludeQuote,
		isIncludeSingleQuote,
		references...)

	return strings.Join(
		slice,
		joiner)
}
