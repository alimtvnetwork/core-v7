package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/namevalue"
)

func VarNameValuesStrings(
	nameValues ...namevalue.Instance,
) []string {
	if len(nameValues) == 0 {
		return []string{}
	}

	items := make([]string, len(nameValues))

	for i, nameVal := range nameValues {
		items[i] = fmt.Sprintf(
			keyValFormat,
			nameVal.Name,
			nameVal.Value)
	}

	return items
}
