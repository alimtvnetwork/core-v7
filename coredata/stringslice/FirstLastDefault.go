package stringslice

import "gitlab.com/auk-go/core/constants"

func FirstLastDefault(slice []string) (first, last string) {
	length := len(slice)

	if length == 0 {
		return constants.EmptyString, constants.EmptyString
	}

	if length == 1 {
		return slice[0], constants.EmptyString
	}

	// length >= 2
	return slice[0], slice[length-1]
}
