package errcore

import "fmt"

func MessageNameValues(
	message string,
	nameValues ...NameVal,
) string {
	if len(nameValues) == 0 {
		return message
	}

	compiledMap := VarNameValues(nameValues...)

	return fmt.Sprintf(
		messageMapFormat,
		message,
		compiledMap)
}
