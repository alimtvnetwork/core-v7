package errcore

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/namevalue"
)

func VarNameValues(
	nameValues ...namevalue.Instance,
) string {
	if len(nameValues) == 0 {
		return ""
	}

	items := VarNameValuesStrings(nameValues...)

	return strings.Join(
		items,
		constants.CommaSpace)
}
