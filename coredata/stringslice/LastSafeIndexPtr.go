package stringslice

import "github.com/alimtvnetwork/core/constants"

// Deprecated: Use LastSafeIndex instead (on non-pointer slice).
func LastSafeIndexPtr(slice []string) int {
	if IsEmptyPtr(slice) {
		return constants.InvalidNotFoundCase
	}

	return len(slice) - constants.One
}
