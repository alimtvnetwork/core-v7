package errcore

import (
	"gitlab.com/auk-go/core/constants"
)

func CombineWithMsgType(
	genericMsg RawErrorType,
	otherMsg string,
	reference interface{},
) string {
	if otherMsg == "" {
		return genericMsg.String() +
			getReferenceMessage(reference)
	}

	return genericMsg.String() +
		constants.Space +
		otherMsg +
		getReferenceMessage(reference)
}
