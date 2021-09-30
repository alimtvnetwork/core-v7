package converters

import "gitlab.com/evatix-go/core/coreappend"

func AnyItemsToStringsSkipOnNil(
	anyItems ...interface{},
) []string {
	return coreappend.PrependAppendAnyItemsToStringsSkipOnNil(
		nil,
		nil,
		anyItems...)
}
