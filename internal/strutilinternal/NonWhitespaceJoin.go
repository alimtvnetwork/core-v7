package strutilinternal

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func NonWhitespaceJoin(slice []string, joiner string) string {
	if slice == nil {
		return constants.EmptyString
	}

	length := len(slice)

	if length == 0 {
		return constants.EmptyString
	}

	return strings.Join(*NonWhitespaceSlicePtr(&slice), joiner)
}
