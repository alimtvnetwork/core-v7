package stringslice

import "gitlab.com/auk-go/core/constants"

func SafeIndexAtUsingLastIndex(
	slice []string,
	lastIndex,
	index int,
) string {
	if lastIndex <= 0 || lastIndex < index || index < 0 {
		return constants.EmptyString
	}

	return slice[index]
}
