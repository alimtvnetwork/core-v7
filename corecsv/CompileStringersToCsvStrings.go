package corecsv

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// CompileStringersToCsvStrings
//
// Formats :
//  - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//  - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//  - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func CompileStringersToCsvStrings(
	isIncludeQuote,
	isIncludeSingleQuote bool, // disable this will give double quote
	compileStringerFunctions ...func() string,
) []string {
	if len(compileStringerFunctions) == 0 {
		return []string{}
	}

	slice := make([]string, len(compileStringerFunctions))

	if isIncludeQuote && isIncludeSingleQuote {
		// single quote
		for i, compilerFunc := range compileStringerFunctions {
			slice[i] = fmt.Sprintf(
				constants.ValueWithSingleQuoteFormat,
				compilerFunc())
		}

		return slice
	} else if isIncludeQuote && !isIncludeSingleQuote {
		// double quote
		for i, compilerFunc := range compileStringerFunctions {
			slice[i] = fmt.Sprintf(
				constants.ValueWithDoubleQuoteFormat,
				compilerFunc())
		}

		return slice
	}

	// no quote
	for i, compilerFunc := range compileStringerFunctions {
		slice[i] = fmt.Sprintf(
			constants.SprintValueFormat,
			compilerFunc())
	}

	return slice
}
