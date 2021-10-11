package errcore

import "fmt"

func VarNameValuesStrings(
	nameValues ...NameVal,
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
