package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// WithCurlyQuotation
//
// Example : {\"%v\"}
func WithCurlyQuotation(
	source interface{},
) string {
	return fmt.Sprintf(
		constants.CurlyQuotationWrapFormat,
		source)
}
