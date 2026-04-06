package stringslice

import "github.com/alimtvnetwork/core/constants"

func LastPtr(slice []string) string {
	return slice[len(slice)-constants.One]
}
