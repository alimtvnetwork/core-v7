package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// WithBracketsQuotation
//
// [\"%v\"]
func WithBracketsQuotation(
	source interface{},
) string {
	return fmt.Sprintf(constants.BracketQuotationWrapFormat, source)
}
