package stringslice

import "github.com/alimtvnetwork/core/constants"

func FirstLastDefaultPtr(slice []string) (first, last string) {
	if len(slice) == 0 {
		return constants.EmptyString, constants.EmptyString
	}

	return FirstLastDefault(slice)
}
