package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// WithParenthesisQuotation
//
// (\"%v\")
func WithParenthesisQuotation(
	source interface{},
) string {
	return fmt.Sprintf(constants.ParenthesisQuotationWrap, source)
}
