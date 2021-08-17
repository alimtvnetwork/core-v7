package utilstringinternal

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToString(any interface{}) string {
	if any == nil {
		return ""
	}

	val := ReflectInterfaceVal(any)

	return fmt.Sprintf(
		constants.SprintValueFormat,
		val)
}
