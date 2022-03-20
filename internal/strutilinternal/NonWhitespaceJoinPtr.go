package strutilinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func NonWhitespaceJoinPtr(slice *[]string, joiner string) string {
	if slice == nil {
		return constants.EmptyString
	}

	length := len(*slice)

	if length == 0 {
		return constants.EmptyString
	}

	return strings.Join(*NonWhitespaceSlicePtr(slice), joiner)
}
