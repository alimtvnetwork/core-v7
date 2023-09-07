package errcore

import (
	"strings"

	"gitlab.com/auk-go/core/namevalue"
)

func VarNameValuesJoiner(
	joiner string,
	nameValues ...namevalue.Instance,
) string {
	if len(nameValues) == 0 {
		return ""
	}

	items := VarNameValuesStrings(nameValues...)

	return strings.Join(
		items,
		joiner)
}
