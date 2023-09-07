package stringslice

import "gitlab.com/auk-go/core/constants"

func LastOrDefaultPtr(slice *[]string) string {
	if slice == nil {
		return constants.EmptyString
	}

	length := len(*slice)

	if length == 0 {
		return constants.EmptyString
	}

	return (*slice)[length-constants.One]
}
