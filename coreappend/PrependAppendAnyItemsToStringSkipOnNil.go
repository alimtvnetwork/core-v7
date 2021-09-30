package coreappend

import "strings"

func PrependAppendAnyItemsToStringSkipOnNil(
	joiner string,
	prependItem, appendItem interface{},
	anyItems ...interface{},
) string {
	slice := PrependAppendAnyItemsToStringsSkipOnNil(
		prependItem,
		appendItem,
		anyItems...)

	return strings.Join(slice, joiner)
}
