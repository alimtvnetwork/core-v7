package strutilinternal

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

// Deprecated: Use NonWhitespaceJoin instead.
func NonWhitespaceJoinPtr(slice []string, joiner string) string {
	if len(slice) == 0 {
		return constants.EmptyString
	}

	return strings.Join(NonWhitespaceSlicePtr(slice), joiner)
}
