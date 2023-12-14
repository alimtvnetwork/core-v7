package corecsv

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// StringsToCsvStrings
//
// Formats :
//   - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//   - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//   - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func StringsToCsvStrings(
	isIncludeQuote,
	isIncludeSingleQuote bool, // disable this will give double quote
	references ...string,
) []string {
	if len(references) == 0 {
		return []string{}
	}

	slice := make([]string, len(references))

	if isIncludeQuote && isIncludeSingleQuote {
		// single quote
		for i, currentReference := range references {
			slice[i] = fmt.Sprintf(
				constants.StringWithSingleQuoteFormat,
				currentReference)
		}

		return slice
	} else if isIncludeQuote && !isIncludeSingleQuote {
		// double quote
		for i, currentReference := range references {
			slice[i] = fmt.Sprintf(
				constants.StringWithDoubleQuoteFormat,
				currentReference)
		}

		return slice
	}

	// no quote
	for i, currentReference := range references {
		slice[i] = fmt.Sprintf(
			constants.SprintStringFormat,
			currentReference)
	}

	return slice
}
