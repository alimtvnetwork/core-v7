package corecsv

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// AnyToTypesCsvStrings
//
// if references empty or len 0 then empty string returned.
//
// Formats :
//   - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//   - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//   - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func AnyToTypesCsvStrings(
	isIncludeQuote,
	isIncludeSingleQuote bool, // disable this will give double quote
	references ...interface{},
) []string {
	if len(references) == 0 {
		return []string{}
	}

	slice := make([]string, len(references))

	if isIncludeQuote && isIncludeSingleQuote {
		// single quote
		for i, input := range references {
			slice[i] = fmt.Sprintf(
				constants.TypeWithSingleQuoteFormat,
				input)
		}

		return slice
	} else if isIncludeQuote && !isIncludeSingleQuote {
		// double quote
		for i, input := range references {
			slice[i] = fmt.Sprintf(
				constants.TypeWithDoubleQuoteFormat,
				input)
		}

		return slice
	}

	// no quote
	for i, input := range references {
		slice[i] = fmt.Sprintf(
			constants.SprintTypeFormat,
			input)
	}

	return slice
}
