package stringslice

import "gitlab.com/auk-go/core/constants"

func LastSafeIndexPtr(slice *[]string) int {
	if IsEmptyPtr(slice) {
		return constants.InvalidNotFoundCase
	}

	return len(*slice) - constants.One
}
