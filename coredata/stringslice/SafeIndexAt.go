package stringslice

import "gitlab.com/auk-go/core/constants"

func SafeIndexAt(slice []string, index int) string {
	if len(slice) == 0 || index < 0 || len(slice)-1 < index {
		return constants.EmptyString
	}

	return slice[index]
}
