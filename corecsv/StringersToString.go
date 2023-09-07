package corecsv

import (
	"fmt"
	"strings"
)

func StringersToString(
	joiner string,
	isIncludeQuote,
	isIncludeSingleQuote bool, // disable this will give double quote
	stringerFunctions ...fmt.Stringer,
) string {
	slice := StringersToCsvStrings(
		isIncludeQuote,
		isIncludeSingleQuote,
		stringerFunctions...)

	return strings.Join(slice, joiner)
}
