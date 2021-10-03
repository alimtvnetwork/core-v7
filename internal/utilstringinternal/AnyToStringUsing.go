package utilstringinternal

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToStringUsing(
	isIncludeFields bool,
	any interface{},
) string {
	if any == nil {
		return ""
	}

	if isIncludeFields {
		return fmt.Sprintf(
			constants.SprintPropertyNameValueFormat,
			any)
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		any)
}
