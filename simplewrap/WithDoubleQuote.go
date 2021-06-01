package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// WithDoubleQuote " + source + " , also take care of any double if available next.
func WithDoubleQuote(source string) string {
	return fmt.Sprintf(
		constants.SprintDoubleQuoteFormat,
		source)
}
