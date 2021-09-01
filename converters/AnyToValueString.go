package converters

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToValueString(any interface{}) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		any)
}
