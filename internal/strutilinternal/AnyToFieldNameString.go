package strutilinternal

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func AnyToFieldNameString(any interface{}) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		any)
}
