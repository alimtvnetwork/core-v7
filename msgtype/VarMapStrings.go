package msgtype

import "fmt"

func VarMapStrings(
	mappedItems map[string]interface{},
) []string {
	if len(mappedItems) == 0 {
		return []string{}
	}

	items := make([]string, len(mappedItems))

	index := 0
	for k, v := range mappedItems {
		items[index] = fmt.Sprintf(
			keyValFormat,
			k,
			v)
		index++
	}

	return items
}
