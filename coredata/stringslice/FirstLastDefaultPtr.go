package stringslice

import "gitlab.com/auk-go/core/constants"

func FirstLastDefaultPtr(slice *[]string) (first, last string) {
	if slice == nil {
		return constants.EmptyString, constants.EmptyString
	}

	return FirstLastDefault(*slice)
}
