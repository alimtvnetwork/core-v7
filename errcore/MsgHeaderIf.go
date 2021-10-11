package errcore

import (
	"fmt"
)

func MsgHeaderIf(
	isHeader bool,
	items ...interface{},
) string {
	if isHeader {
		return MsgHeader(items...)
	}

	return fmt.Sprint(items...)
}
