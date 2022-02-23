package enumimpl

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

func NameWithValue(
	value interface{},
) string {
	return fmt.Sprintf(
		constants.EnumNameValueFormat,
		value,
		value)
}
