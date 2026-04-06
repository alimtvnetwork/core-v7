package stringslice

import "github.com/alimtvnetwork/core/constants"

func LastSafeIndexPtr(slice []string) int {
	if IsEmptyPtr(slice) {
		return constants.InvalidNotFoundCase
	}

	return len(slice) - constants.One
}
