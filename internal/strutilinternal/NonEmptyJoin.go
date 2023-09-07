package strutilinternal

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func NonEmptyJoin(slice *[]string, joiner string) string {
	if slice == nil {
		return constants.EmptyString
	}

	length := len(*slice)

	if length == 0 {
		return constants.EmptyString
	}

	return strings.Join(*NonEmptySlicePtr(slice), joiner)
}
