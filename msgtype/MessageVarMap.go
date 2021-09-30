package msgtype

import "fmt"

func MessageVarMap(
	message string,
	mappedItems map[string]interface{},
) string {
	if len(mappedItems) == 0 {
		return message
	}

	compiledMap := VarMap(mappedItems)

	return fmt.Sprintf(
		messageMapFormat,
		message,
		compiledMap)
}
