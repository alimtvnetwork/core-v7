package coreappend

import "strings"

func AppendAnyItemsToStringSkipOnNil(
	joiner string,
	appendItem interface{},
	anyItems ...interface{},
) string {
	slice := PrependAppendAnyItemsToStringsSkipOnNil(
		nil,
		appendItem,
		anyItems...)

	return strings.Join(slice, joiner)
}
