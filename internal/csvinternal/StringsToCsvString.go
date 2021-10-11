package csvinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// StringsToCsvString
//
// if references empty or len 0 then empty string returned.
//
// Final join whole lines with the joiner given (... joiner item)
//
// Formats :
//  - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//  - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//  - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func StringsToCsvString(
	joiner string,
	isIncludeQuote,
	isIncludeSingleQuote bool,
	references ...string,
) string {
	if len(references) == 0 {
		return constants.EmptyString
	}

	slice := StringsToCsvStrings(
		isIncludeQuote,
		isIncludeSingleQuote,
		references...)

	return strings.Join(
		slice,
		joiner)
}
