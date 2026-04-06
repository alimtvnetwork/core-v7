package stringslice

import "github.com/alimtvnetwork/core/constants"

func LastOrDefaultPtr(slice []string) string {
	length := len(slice)

	if length == 0 {
		return constants.EmptyString
	}

	return slice[length-constants.One]
}
