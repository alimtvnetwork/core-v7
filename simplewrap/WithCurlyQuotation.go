package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// WithCurlyQuotation
//
// {\"%v\"}
func WithCurlyQuotation(
	source interface{},
) string {
	return fmt.Sprintf(constants.CurlyQuotationWrap, source)
}
