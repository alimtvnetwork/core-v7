package errcore

import "fmt"

func getReferenceMessage(
	reference interface{},
) string {
	if reference == nil {
		return ""
	}

	currentString, isString := reference.(string)
	if isString && currentString == "" {
		return ""
	}

	return fmt.Sprintf(
		ReferenceFormat,
		reference)
}
