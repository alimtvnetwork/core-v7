package stringslice

import "gitlab.com/evatix-go/core/constants"

func SafeIndexAtUsingLastIndex(
	slice []string,
	lastIndex,
	index int,
) string {
	if lastIndex == 0 || lastIndex < index {
		return constants.EmptyString
	}

	return slice[index]
}
