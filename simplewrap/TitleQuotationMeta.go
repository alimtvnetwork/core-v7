package simplewrap

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// TitleQuotationMeta
//
// Example :
//  - constants.QuotationTitleMetaWrapFormat
//  -  "%v: \"%v\" (%v)"
func TitleQuotationMeta(
	title,
	value,
	meta interface{},
) string {
	return fmt.Sprintf(
		constants.QuotationTitleMetaWrapFormat,
		title,
		value,
		meta)
}
