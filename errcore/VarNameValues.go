package errcore

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func VarNameValues(
	nameValues ...NameVal,
) string {
	if len(nameValues) == 0 {
		return ""
	}

	items := VarNameValuesStrings(nameValues...)

	return strings.Join(
		items,
		constants.CommaSpace)
}
