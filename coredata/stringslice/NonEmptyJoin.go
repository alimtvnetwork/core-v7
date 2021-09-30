package stringslice

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func NonEmptyJoin(slice []string, joiner string) string {
	if slice == nil {
		return constants.EmptyString
	}

	length := len(slice)

	if length == 0 {
		return constants.EmptyString
	}

	return strings.Join(*NonEmptySlicePtr(&slice), joiner)
}
