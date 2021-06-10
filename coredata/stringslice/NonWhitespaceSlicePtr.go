package stringslice

import (
	"gitlab.com/evatix-go/core/internal/stringutil"
)

func NonWhitespaceSlicePtr(slice *[]string) *[]string {
	if slice == nil {
		return &[]string{}
	}

	length := len(*slice)

	if length == 0 {
		return &[]string{}
	}

	newSlice := MakeDefault(length)

	for _, s := range *slice {
		if stringutil.IsEmptyOrWhitespace(s) {
			continue
		}

		newSlice = append(newSlice, s)
	}

	return &newSlice
}
