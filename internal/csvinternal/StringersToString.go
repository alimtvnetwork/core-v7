package csvinternal

import (
	"fmt"
	"strings"
)

func StringersToString(
	joiner string,
	isIncludeQuote,
	isIncludeSingleQuote bool,
	stringerFunctions ...fmt.Stringer,
) string {
	slice := StringersToCsvStrings(
		isIncludeQuote,
		isIncludeSingleQuote,
		stringerFunctions...)

	return strings.Join(slice, joiner)
}
