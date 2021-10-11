package errcore

import "strings"

func VarNameValuesJoiner(
	joiner string,
	nameValues ...NameVal,
) string {
	if len(nameValues) == 0 {
		return ""
	}

	items := VarNameValuesStrings(nameValues...)

	return strings.Join(
		items,
		joiner)
}
