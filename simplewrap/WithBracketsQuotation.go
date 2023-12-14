package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

// WithBracketsQuotation
//
// [\"%v\"]
func WithBracketsQuotation(
	source interface{},
) string {
	return fmt.Sprintf(
		constants.BracketQuotationWrapFormat,
		toString(source))
}
