package stringslice

import "gitlab.com/auk-go/core/constants"

func FirstOrDefault(slice []string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return slice[0]
}
