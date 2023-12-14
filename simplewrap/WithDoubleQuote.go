package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// WithDoubleQuote " + source + " , also take care of any double if available next.
func WithDoubleQuote(source string) string {
	return fmt.Sprintf(
		constants.SprintStartStringEndCharFormat,
		constants.DoubleQuoteChar,
		source,
		constants.DoubleQuoteChar)
}
