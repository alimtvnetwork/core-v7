package corestr

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func AnyToString(
	isIncludeFieldName bool,
	any interface{},
) string {
	if any == "" {
		return constants.EmptyString
	}

	val := reflectInterfaceVal(any)

	if isIncludeFieldName {
		return fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			val)
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		val)
}
