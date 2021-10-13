package csvinternal

import "strings"

func CompileStringersToString(
	joiner string,
	isIncludeQuote,
	isIncludeSingleQuote bool,
	compileStringerFunctions ...func() string,
) string {
	slice := CompileStringersToCsvStrings(
		isIncludeQuote,
		isIncludeSingleQuote,
		compileStringerFunctions...)

	return strings.Join(slice, joiner)
}
