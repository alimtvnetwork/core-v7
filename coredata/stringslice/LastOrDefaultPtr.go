package stringslice

import "github.com/alimtvnetwork/core/constants"

// Deprecated: Use LastOrDefault instead.
func LastOrDefaultPtr(slice []string) string {
	length := len(slice)

	if length == 0 {
		return constants.EmptyString
	}

	return slice[length-constants.One]
}
