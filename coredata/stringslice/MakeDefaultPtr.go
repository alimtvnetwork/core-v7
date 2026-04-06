package stringslice

import "github.com/alimtvnetwork/core/constants"

func MakeDefaultPtr(capacity int) []string {
	return make([]string, constants.Zero, capacity)
}
