package errcore

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// StringLinesToQuoteLinesToSingle
//
// Each line will be wrapped with "\"%s\", quotation and comma
func StringLinesToQuoteLinesToSingle(lines []string) string {
	slice := StringLinesToQuoteLines(lines)

	return strings.Join(
		slice,
		constants.NewLineUnix)
}
