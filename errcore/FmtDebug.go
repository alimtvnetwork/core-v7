package errcore

import "fmt"

func FmtDebug(
	format string,
	items ...interface{},
) {
	fmt.Printf(format, items...)
}

func ValidPrint(
	isValid bool,
	items ...interface{},
) {
	if isValid {
		fmt.Print(items...)
	}
}

func FailedPrint(
	isFailed bool,
	items ...interface{},
) {
	if isFailed {
		fmt.Print(items...)
	}
}
