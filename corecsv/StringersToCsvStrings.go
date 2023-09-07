package corecsv

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// StringersToCsvStrings
//
// Formats :
//  - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//  - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//  - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func StringersToCsvStrings(
	isIncludeQuote,
	isIncludeSingleQuote bool, // disable this will give double quote
	stringerFunctions ...fmt.Stringer,
) []string {
	if len(stringerFunctions) == 0 {
		return []string{}
	}

	slice := make([]string, len(stringerFunctions))

	if isIncludeQuote && isIncludeSingleQuote {
		// single quote
		for i, stringerFunc := range stringerFunctions {
			slice[i] = fmt.Sprintf(
				constants.ValueWithSingleQuoteFormat,
				stringerFunc.String())
		}

		return slice
	} else if isIncludeQuote && !isIncludeSingleQuote {
		// double quote
		for i, stringerFunc := range stringerFunctions {
			slice[i] = fmt.Sprintf(
				constants.ValueWithDoubleQuoteFormat,
				stringerFunc.String())
		}

		return slice
	}

	// no quote
	for i, stringerFunc := range stringerFunctions {
		slice[i] = fmt.Sprintf(
			constants.SprintValueFormat,
			stringerFunc.String())
	}

	return slice
}
