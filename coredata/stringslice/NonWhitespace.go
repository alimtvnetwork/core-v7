package stringslice

import (
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

// NonWhitespace
//
// Don't include line which is empty or whitespace.
func NonWhitespace(
	slice []string,
) []string {
	if slice == nil {
		return []string{}
	}

	length := len(slice)

	if length == 0 {
		return []string{}
	}

	newSlice := MakeDefault(length)

	for _, s := range slice {
		if strutilinternal.IsEmptyOrWhitespace(s) {
			continue
		}

		newSlice = append(newSlice, s)
	}

	return newSlice
}
