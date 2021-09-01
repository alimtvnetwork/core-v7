package converters

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToFullNameValueString(any interface{}) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintFullPropertyNameValueFormat,
		any)
}
