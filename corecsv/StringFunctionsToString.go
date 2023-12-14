package corecsv

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// StringFunctionsToString
//
// Formats :
//   - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//   - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//   - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func StringFunctionsToString(
	isIncludeQuote,
	isIncludeSingleQuote bool, // disable this will give double quote
	stringerFunctions ...func() string,
) []string {
	if len(stringerFunctions) == 0 {
		return []string{}
	}

	slice := make([]string, len(stringerFunctions))

	if isIncludeQuote && isIncludeSingleQuote {
		// single quote
		for i, stringerFunc := range stringerFunctions {
			slice[i] = fmt.Sprintf(
				constants.StringWithSingleQuoteFormat,
				stringerFunc())
		}

		return slice
	} else if isIncludeQuote && !isIncludeSingleQuote {
		// double quote
		for i, stringerFunc := range stringerFunctions {
			slice[i] = fmt.Sprintf(
				constants.StringWithDoubleQuoteFormat,
				stringerFunc())
		}

		return slice
	}

	// no quote
	for i, stringerFunc := range stringerFunctions {
		slice[i] = fmt.Sprintf(
			constants.SprintStringFormat,
			stringerFunc())
	}

	return slice
}
