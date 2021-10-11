package errcore

import "fmt"

func FmtDebugIf(
	isDebug bool,
	format string,
	items ...interface{},
) {
	if !isDebug {
		return
	}

	fmt.Printf(format, items...)
}
