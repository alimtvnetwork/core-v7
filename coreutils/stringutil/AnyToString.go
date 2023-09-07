package stringutil

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func AnyToString(any interface{}) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(constants.SprintValueFormat, any)
}
