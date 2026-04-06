package stringslice

import "github.com/alimtvnetwork/core/constants"

// Deprecated: Use FirstOrDefault instead.
func FirstOrDefaultPtr(slice []string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return slice[0]
}
