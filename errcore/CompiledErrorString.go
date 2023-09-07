package errcore

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

func CompiledErrorString(mainErr error, additionalMessage string) string {
	if mainErr == nil {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		constants.MessageWrapMessageFormat,
		mainErr.Error(),
		additionalMessage)
}
