package reqtype

import "strings"

func RangesString(
	joiner string,
	requests ...Request,
) string {
	slice := RangesStrings(requests...)

	return strings.Join(slice, joiner)
}
