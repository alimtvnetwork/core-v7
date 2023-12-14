package simplewrap

import (
	"fmt"
	
	"gitlab.com/auk-go/core/constants"
)

// TitleQuotationMeta
//
// Example :
//   - constants.QuotationTitleMetaWrapFormat
//   - "%v: \"%v\" (%v)"
func TitleQuotationMeta(
	title,
	value,
	meta interface{},
) string {
	return fmt.Sprintf(
		constants.QuotationTitleMetaWrapFormat,
		toString(title),
		toString(value),
		meta)
}
