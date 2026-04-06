package stringslice

import "github.com/alimtvnetwork/core/constants"

// Deprecated: Use MakeDefault instead.
func MakeDefaultPtr(capacity int) []string {
	return make([]string, constants.Zero, capacity)
}
