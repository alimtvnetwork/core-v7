package errcore

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

func VarMap(
	mappedItems map[string]interface{},
) string {
	if len(mappedItems) == 0 {
		return ""
	}

	items := VarMapStrings(mappedItems)

	return strings.Join(
		items,
		constants.CommaSpace)
}
