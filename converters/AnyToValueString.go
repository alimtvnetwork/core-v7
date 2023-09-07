package converters

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func AnyToValueString(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem)
}
