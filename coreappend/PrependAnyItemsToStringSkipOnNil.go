package coreappend

import "strings"

func PrependAnyItemsToStringSkipOnNil(
	joiner string,
	prependItem interface{},
	anyItems ...interface{},
) string {
	slice := PrependAppendAnyItemsToStringsSkipOnNil(
		prependItem,
		nil,
		anyItems...)

	return strings.Join(slice, joiner)
}
