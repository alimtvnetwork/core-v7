package stringslice

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

func NonEmptyJoinPtr(slice []string, joiner string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return strings.Join(NonEmptySlicePtr(slice), joiner)
}
