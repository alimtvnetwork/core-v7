package errcore

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// PathMeaningfulMessage skip error if messages empty or length 0
func PathMeaningfulMessage(
	rawErrType RawErrorType,
	funcName string,
	location string,
	messages ...string,
) error {
	if len(messages) == 0 {
		return nil
	}

	messagesCompiled := strings.Join(messages, constants.Space)
	errMsg := "location: [" + location + "], " + messagesCompiled

	return rawErrType.Error(
		funcName,
		errMsg)
}
