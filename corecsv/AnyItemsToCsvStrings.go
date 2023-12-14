package corecsv

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// AnyItemsToCsvStrings
//
// Formats :
//   - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//   - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//   - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func AnyItemsToCsvStrings(
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
		for i, currentReference := range references {
			slice[i] = fmt.Sprintf(
				constants.StringWithSingleQuoteFormat,
				toString(currentReference))
		}

		return slice
	} else if isIncludeQuote && !isIncludeSingleQuote {
		// double quote
		for i, currentReference := range references {
			slice[i] = fmt.Sprintf(
				constants.StringWithDoubleQuoteFormat,
				toString(currentReference))
		}

		return slice
	}

	// no quote
	for i, currentReference := range references {
		slice[i] = fmt.Sprintf(
			constants.SprintStringFormat,
			toString(currentReference))
	}

	return slice
}
