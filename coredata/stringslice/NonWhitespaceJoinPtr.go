package stringslice

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

func NonWhitespaceJoinPtr(slice []string, joiner string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return strings.Join(NonWhitespacePtr(slice), joiner)
}
