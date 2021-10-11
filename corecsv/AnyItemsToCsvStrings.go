package corecsv

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// AnyItemsToCsvStrings
//
// Formats :
//  - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//  - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//  - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func AnyItemsToCsvStrings(
	isIncludeQuote,
	isIncludeSingleQuote bool,
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
				constants.ValueWithSingleQuoteFormat,
				currentReference)
		}

		return slice
	} else if isIncludeQuote && !isIncludeSingleQuote {
		// double quote
		for i, currentReference := range references {
			slice[i] = fmt.Sprintf(
				constants.ValueWithDoubleQuoteFormat,
				currentReference)
		}

		return slice
	}

	// no quote
	for i, currentReference := range references {
		slice[i] = fmt.Sprintf(
			constants.SprintValueFormat,
			currentReference)
	}

	return slice
}
