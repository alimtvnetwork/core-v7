package stringslice

import "gitlab.com/auk-go/core/constants"

func LastIndexPtr(slice *[]string) int {
	return len(*slice) - constants.One
}
