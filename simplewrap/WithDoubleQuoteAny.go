package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

// WithDoubleQuoteAny
//
//	Alias for ToJsonName
//
//	" + source + " , also take care of any double if available next.
func WithDoubleQuoteAny(source interface{}) string {
	return fmt.Sprintf(
		constants.SprintValueDoubleQuotationFormat,
		toString(source))
}
