package converters

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToValueStringWithType(any interface{}) string {
	if any == nil {
		return fmt.Sprintf(
			constants.SprintNilValueTypeInParentThesisFormat,
			any)
	}

	return fmt.Sprintf(
		constants.SprintValueWithTypeFormat,
		any,
		any)
}
