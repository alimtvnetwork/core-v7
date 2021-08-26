package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// WithDoubleQuoteAny " + source + " , also take care of any double if available next.
func WithDoubleQuoteAny(source interface{}) string {
	return fmt.Sprintf(
		constants.SprintValueDoubleQuotationFormat,
		source)
}
