package converters

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func AnyToAnyItemsThenJoin(
	isSkipOnNil bool,
	joiner string,
	anySlice interface{},
) string {
	if anySlice == nil {
		return constants.EmptyString
	}

	anyStrings := AnyToStrings(isSkipOnNil, anySlice)

	return strings.Join(anyStrings, joiner)
}
