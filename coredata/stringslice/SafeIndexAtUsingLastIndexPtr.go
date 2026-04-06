package stringslice

import "github.com/alimtvnetwork/core/constants"

// Deprecated: Use SafeIndexAtUsingLastIndex instead.
func SafeIndexAtUsingLastIndexPtr(
	slice []string,
	lastIndex,
	index int,
) string {
	if lastIndex == 0 || lastIndex < 0 || lastIndex < index || index < 0 {
		return constants.EmptyString
	}

	return slice[index]
}
