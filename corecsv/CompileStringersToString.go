package corecsv

import "strings"

func CompileStringersToString(
	joiner string,
	isIncludeQuote,
	isIncludeSingleQuote bool, // disable this will give double quote
	compileStringerFunctions ...func() string,
) string {
	slice := CompileStringersToCsvStrings(
		isIncludeQuote,
		isIncludeSingleQuote,
		compileStringerFunctions...)

	return strings.Join(slice, joiner)
}
