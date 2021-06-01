package converters

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToString(any interface{}) string {
	if any == nil {
		return ""
	}

	return fmt.Sprintf(constants.SprintValueFormat, any)
}
