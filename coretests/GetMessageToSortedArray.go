package coretests

import (
	"fmt"
)

func GetMessageToSortedArray(
	isPrint bool,
	isSort bool,
	message string,
) []string {
	if isPrint {
		fmt.Println(message)
	}

	return GetTrimmedNonEmptySpaceSplit(
		message,
		isSort)
}
