package stringutil

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func AnyToStringNameField(any interface{}) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(constants.SprintPropertyNameValueFormat, any)
}
