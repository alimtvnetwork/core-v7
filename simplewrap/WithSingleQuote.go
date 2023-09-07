package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// WithSingleQuote ' + source + ' , also take care of any double if available next.
func WithSingleQuote(source string) string {
	return fmt.Sprintf(
		constants.SprintSingleQuoteFormat,
		source)
}
