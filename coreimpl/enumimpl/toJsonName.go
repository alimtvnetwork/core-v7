package enumimpl

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// toJsonName
//
//	" + source + " , also take care of any double if available next.
func toJsonName(source interface{}) string {
	return fmt.Sprintf(
		constants.SprintValueDoubleQuotationFormat,
		source)
}
