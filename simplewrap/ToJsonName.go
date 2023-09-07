package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// ToJsonName
//
//  Alias for WithDoubleQuoteAny
//
//  " + source + " , also take care of any double if available next.
func ToJsonName(source interface{}) string {
	return fmt.Sprintf(
		constants.SprintValueDoubleQuotationFormat,
		source)
}
