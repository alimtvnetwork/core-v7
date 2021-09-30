package converters

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func AnyItemsJoin(
	joiner string,
	anyItems ...interface{},
) string {
	if anyItems == nil {
		return constants.EmptyString
	}

	anyStrings := AnyItemsToStringsSkipOnNil(anyItems...)

	return strings.Join(anyStrings, joiner)
}
