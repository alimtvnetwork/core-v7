package errcore

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// PathMeaningFulMessage skip error if messages empty or length 0
func PathMeaningFulMessage(
	msgType Variation,
	funcName string,
	location string,
	messages ...string,
) error {
	if len(messages) == 0 {
		return nil
	}

	messagesCompiled := strings.Join(messages, constants.Space)
	errMsg := "location: [" + location + "], " + messagesCompiled

	return msgType.Error(
		funcName,
		errMsg)
}
