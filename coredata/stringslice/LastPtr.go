package stringslice

import "github.com/alimtvnetwork/core/constants"

// Deprecated: Use Last instead.
func LastPtr(slice []string) string {
	return slice[len(slice)-constants.One]
}
