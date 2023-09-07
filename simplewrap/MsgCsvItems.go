package simplewrap

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/csvinternal"
)

func MsgCsvItems(
	msg string,
	csvItems ...interface{},
) string {
	csvString := csvinternal.AnyItemsToStringDefault(
		csvItems...)

	return fmt.Sprintf(
		constants.ValueWrapValueFormat,
		msg,
		csvString)
}
