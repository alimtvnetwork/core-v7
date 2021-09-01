package converters

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToString(
	isIncludeFullName bool,
	any interface{},
) string {
	if any == nil {
		return ""
	}

	if isIncludeFullName {
		return fmt.Sprintf(
			constants.SprintFullPropertyNameValueFormat,
			any)
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		any)
}
