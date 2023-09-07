package stringslice

import "gitlab.com/auk-go/core/constants"

func LastPtr(slice *[]string) string {
	return (*slice)[len(*slice)-constants.One]
}
