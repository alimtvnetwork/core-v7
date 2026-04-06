package stringslice

import "github.com/alimtvnetwork/core/constants"

func FirstOrDefaultPtr(slice []string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return slice[0]
}
